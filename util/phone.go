package util

import (
	config "Raising/conf"
	"errors"
	"sync"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	aliyunUtil "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/patrickmn/go-cache"
)

type aliyun struct {
	verificationCodeCache    *cache.Cache // 验证码 5 分钟过期
	verificationCodeReqCache *cache.Cache // 一分钟内只能发送一次验证码
}

var (
	aliyunOnce   sync.Once
	aliyunEntity *aliyun
)

func getAliyunEntity() *aliyun {
	aliyunOnce.Do(func() {
		aliyunEntity = new(aliyun)
		aliyunEntity.verificationCodeReqCache = cache.New(time.Minute, time.Minute)
		aliyunEntity.verificationCodeCache = cache.New(time.Minute*5, time.Minute*5)
	})
	return aliyunEntity

}
func (ali *aliyun) SendVerificationCode(phoneNumber string) (err error) {
	// 验证是否可以获取验证码（1分钟有效期）
	_, found := ali.verificationCodeReqCache.Get(phoneNumber)
	if found {
		err = errors.New("请勿重复发送验证码")
		return
	}

	// 生成验证码
	verifyCode := CreateRandCode()

	// 发送短信
	err = ali.SendSms(ali.getVerifyCodeReq(phoneNumber, verifyCode))
	if err != nil {
		return
	}

	// 验证码加入缓存
	ali.verificationCodeReqCache.SetDefault(phoneNumber, 1)
	ali.verificationCodeCache.SetDefault(phoneNumber, verifyCode)

	return
}
func (ali *aliyun) CheckVerificationCode(phoneNumber, verificationCode string) (err error) {
	cacheCode, found := ali.verificationCodeCache.Get(phoneNumber)
	if !found {
		err = errors.New("验证码已失效")
		return
	}

	cc, sure := cacheCode.(string)
	if !sure {
		err = errors.New("内部服务出错")
		return
	}
	if cc != verificationCode {
		err = errors.New("验证码输入错误")
		return
	}
	return
}
func (*aliyun) CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}
func (ali *aliyun) SendSms(req dysmsapi20170525.SendSmsRequest) (_err error) {
	// TODO your key，from config
	client, _err := ali.CreateClient(tea.String(config.AliAccessKey), tea.String(config.AliSerectKey))
	if _err != nil {
		return _err
	}

	defer func() {
		if r := tea.Recover(recover()); r != nil {
			_err = r
		}
	}()

	runtime := &aliyunUtil.RuntimeOptions{}
	result, _err := client.SendSmsWithOptions(&req, runtime)
	if _err != nil {
		return _err
	}

	if *result.Body.Code != "OK" {
		_err = errors.New(result.String())
		return
	}

	return _err
}
func (ali *aliyun) getVerifyCodeReq(phoneNumber, code string) (req dysmsapi20170525.SendSmsRequest) {
	// TODO SignName TemplateCode
	req = dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(config.SigName),
		TemplateCode:  tea.String(config.TemplateCode),
		PhoneNumbers:  tea.String(phoneNumber),
		TemplateParam: tea.String(`{"code":"` + code + `"}`),
	}
	return
}
