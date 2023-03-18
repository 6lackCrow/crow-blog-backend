package encryptUtil

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptPassword), nil
}

// EqualsPassword 对比密码是否正确
func EqualsPassword(hashedPassword, password string) bool {
	// 使用 bcrypt 当中的 CompareHashAndPassword 对比密码是否正确，第一个参数为加密后的密码，第二个参数为未加密的密码
	// 对比密码是否正确会返回一个异常，按照官方的说法是只要异常是 nil 就证明密码正确
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
