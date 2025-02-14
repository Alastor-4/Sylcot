package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			c.Abort()
			return
		}

		// // Validar el token JWT
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	return []byte("tu_clave_secreta"), nil
		// })
		// if err != nil || !token.Valid {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		// 	c.Abort()
		// 	return
		// }

		// claims := token.Claims.(jwt.MapClaims)
		// userID := uint(claims["userID"].(float64))

		// // Verificar si el usuario está verificado
		// var user models.User
		// if err := database.DB.First(&user, userID).Error; err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		// 	c.Abort()
		// 	return
		// }

		// if !user.IsVerified {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "Usuario no verificado"})
		// 	c.Abort()
		// 	return
		// }

		// // Guardar el ID del usuario en el contexto
		// c.Set("userID", userID)
		// c.Next()
	}
}

func SendVerificationEmail(c *gin.Context) {
	// var user models.User
	// email := c.Query("email") // Obtener el correo desde la solicitud

	// // Buscar al usuario por correo
	// if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	// 	return
	// }

	// // Generar un token único
	// token := uuid.New().String()
	// user.Token = token
	// database.DB.Save(&user)

	// // Enviar el correo
	// m := gomail.NewMessage()
	// m.SetHeader("From", "tusistema@example.com")
	// m.SetHeader("To", user.Email)
	// m.SetHeader("Subject", "Verifica tu correo electrónico")
	// m.SetBody("text/html", `<p>Haz clic <a href="http://localhost:8080/auth/verify?token=`+token+`">aquí</a> para verificar tu correo.</p>`)

	// d := gomail.NewDialer("smtp.example.com", 587, "user", "password") // Configura tu SMTP
	// if err := d.DialAndSend(m); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enviando el correo"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Correo de verificación enviado"})
}

func Register(c *gin.Context) {
}

func Login(c *gin.Context) {
}

func VerifyEmail(c *gin.Context) {
	// token := c.Query("token") // Obtener el token desde la URL

	// var user models.User
	// if err := database.DB.Where("token = ?", token).First(&user).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Token inválido"})
	// 	return
	// }

	// // Activar la cuenta
	// user.IsVerified = true
	// user.Token = "" // Limpiar el token
	// database.DB.Save(&user)

	// Redirigir a la aplicación
	c.Redirect(http.StatusMovedPermanently, "http://tuaplicacion.com/login")
}
