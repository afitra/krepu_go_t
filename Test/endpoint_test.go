package test

import (
	"bytes"
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

var (
	baseUrl    = "http://localhost:3000/"
	baseToken  = ""
	user_name  = ""
	password   = ""
	no_kontrak = ""
)

func setBaseToken(token string) {
	baseToken = token
}

func setValue(variable interface{}, newValue interface{}) {
	reflect.ValueOf(variable).Elem().Set(reflect.ValueOf(newValue))
}

func generateRandomNik() string {
	var nik string
	err := faker.FakeData(&nik)
	if err != nil {
		panic(err)
	}
	return nik
}

func generateRandomUserName() string {
	var userName string
	err := faker.FakeData(&userName)
	if err != nil {
		panic(err)
	}
	return userName
}

func createRequest(method string, url string, requestBody map[string]interface{}) (newMap map[string]interface{}, err error) {

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, baseUrl+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, err
}

func createRequestWithToken(method, url string, requestBody map[string]interface{}) (map[string]interface{}, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, baseUrl+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", baseToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, err
}

func TestRegisterUser(t *testing.T) {
	randomNik := generateRandomNik()
	randomUserName := generateRandomUserName()

	setValue(&user_name, randomUserName)
	setValue(&password, "123456")

	requestBody := map[string]interface{}{
		"nik":           randomNik,
		"full_name":     "Afitra Mamor Bikhoir",
		"legal_name":    "Afitra Mamor Bikhoir",
		"tempat_lahir":  "kediri",
		"tanggal_lahir": "03-03-1996",
		"gaji":          16000000,
		"foto_ktp":      "https://img.inews.co.id/media/600/files/networks/2022/06/20/ffb4c_anya-geraldine.jpg",
		"foto_selfie":   "https://img.inews.co.id/media/600/files/networks/2022/06/20/ffb4c_anya-geraldine.jpg",
		"user_name":     randomUserName,
		"password":      "123456",
	}

	response, err := createRequest(http.MethodPost, "public/user/register", requestBody)
	if err != nil {
		t.Errorf("Terjadi kesalahan, %s", err.Error())
	}

	assert.NotEmpty(t, response["code"])
	assert.NotEmpty(t, response["status"])
	assert.NotEmpty(t, response["message"])

	assert.Equal(t, response["code"], "201")
	assert.Equal(t, response["status"], "success")
	assert.Equal(t, response["message"], "Data Sedang di proses")
}

func TestLogin(t *testing.T) {

	requestBody := map[string]interface{}{
		"user_name": user_name,
		"password":  password,
	}

	response, err := createRequest(http.MethodPost, "public/user/login", requestBody)
	if err != nil {
		t.Errorf("Terjadi kesalahan,  %s", err.Error())
	}

	data, _ := response["data"].(map[string]interface{})

	assert.NotEmpty(t, response["code"])
	assert.NotEmpty(t, response["status"])
	assert.NotEmpty(t, response["message"])
	assert.NotEmpty(t, response["data"])

	assert.Equal(t, response["code"], "200")
	assert.Equal(t, response["status"], "success")
	assert.Equal(t, response["message"], "Data Sedang di proses")

	token, ok := data["token"].(string)
	assert.True(t, ok, "Tidak dapat mendapatkan token dari respons")

	assert.NotEmpty(t, token, "Data tidak berisi token")
	assert.IsType(t, "", token, "Tipe data token tidak sesuai, diharapkan string")

	setValue(&baseToken, data["token"])

}

func TestTransactionInquiry(t *testing.T) {

	requestBody := map[string]interface{}{
		"otr":        1000,
		"admin_fee":  1000,
		"cicilan":    1000,
		"bunga":      2,
		"nama_asset": "rumah",
		"tenor":      1,
		"pengajuan":  100000,
	}

	response, err := createRequestWithToken(http.MethodPost, "private/transaction/inquiry", requestBody)
	if err != nil {
		t.Errorf("Terjadi kesalahan, %s", err.Error())
	}
	data, _ := response["data"].(map[string]interface{})

	assert.NotEmpty(t, response["code"])
	assert.NotEmpty(t, response["status"])
	assert.NotEmpty(t, response["message"])

	assert.Equal(t, response["code"], "200")
	assert.Equal(t, response["status"], "success")
	assert.Equal(t, response["message"], "Data Sedang di proses")

	no_kontrak, ok := data["no_kontrak"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, no_kontrak)
	assert.IsType(t, "", no_kontrak, "Tipe data token tidak sesuai, diharapkan string")

	setValue(&no_kontrak, no_kontrak)
}

func TestAdminTransactionPay(t *testing.T) {

	requestBody := map[string]interface{}{
		"no_kontrak": no_kontrak,
	}

	response, err := createRequestWithToken(http.MethodPost, "admin/transaction/pay", requestBody)
	if err != nil {
		t.Errorf("Terjadi kesalahan, %s", err.Error())
	}

	assert.NotEmpty(t, response["code"])
	assert.NotEmpty(t, response["status"])
	assert.NotEmpty(t, response["message"])

	assert.Equal(t, response["code"], "401")
	assert.Equal(t, response["status"], "failed")
	assert.Equal(t, response["message"], "you dont have access this feature")

}
