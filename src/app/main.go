package main

import (
	"net/http"

	"reflect"

	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "my JSON message",
			"status":   http.StatusOK,
			"computed": computeMath(5)})
	})

	r.GET("/getXML", func(c *gin.Context) {
		number, _ := computeMathAny(13)
		c.XML(http.StatusOK, gin.H{
			"message":  "my XML message",
			"status":   http.StatusOK,
			"computed": number})
	})

	r.GET("/getJSONnumber/:number", func(c *gin.Context) {
		number := c.Param("number")
		_, stringNumber := computeMathAny(number)
		c.JSON(http.StatusOK, gin.H{
			"message":  "my JSON + number message",
			"status":   http.StatusOK,
			"computed": stringNumber})
	})

	r.GET("/getJSONnumberextra/:number/:extranumber", func(c *gin.Context) {
		number := c.Param("number")
		extranumber := c.Param("extranumber")
		_, stringNumber := computeMathAny(number)
		_, stringNumber2 := computeMathAny2(extranumber)
		c.JSON(http.StatusOK, gin.H{
			"message":   "my JSON + number message",
			"status":    http.StatusOK,
			"computed":  stringNumber,
			"computed2": stringNumber2})
	})

	r.GET("/getXMLnumber/:number", func(c *gin.Context) {
		number := c.Param("number")
		_, stringNumber := computeMathAny(number)
		c.XML(http.StatusOK, gin.H{
			"message":  "my XML + number message",
			"status":   http.StatusOK,
			"computed": stringNumber})
	})

	r.GET("/getXMLnumberextra/:number/:extranumber", func(c *gin.Context) {
		number := c.Param("number")
		extranumber := c.Param("extranumber")
		_, stringNumber := computeMathAny(number)
		_, stringNumber2 := computeMathAny2(extranumber)
		c.XML(http.StatusOK, gin.H{
			"message":   "my XML + number message",
			"status":    http.StatusOK,
			"computed":  stringNumber,
			"computed2": stringNumber2})
	})

	r.GET("/getString/:number", func(c *gin.Context) {
		number := c.Param("number")
		c.String(http.StatusOK, reflect.TypeOf(number).String())
	})

	r.GET("getStringNumber/:number", func(c *gin.Context) {
		number := c.Param("number")
		_, stringNumber := computeMathAny(number)
		c.String(http.StatusOK, stringNumber)
	})

	r.GET("getStringNumberInterface/:number", func(c *gin.Context) {
		number := c.Param("number")
		c.String(http.StatusOK, computeMathAnyInterface(number).(string))
	})

	r.Run(":3000")
}

func computeMath(input int) (output int) {
	output = 3 * input
	return
}

func computeMathAny(input interface{}) (outputInt int, outputString string) {
	switch t := input.(type) {
	case int:
		outputInt = 3 * t
		outputString = strconv.Itoa(outputInt)
	case string:
		temp, _ := strconv.Atoi(t)
		outputInt = 3 * temp
		outputString = strconv.Itoa(outputInt)
	default:
		outputInt = 99
		outputString = "99"
	}
	return
}

func computeMathAny2(input interface{}) (outputInt int, outputString string) {
	switch t := input.(type) {
	case int:
		outputInt = 7 * t
		outputString = strconv.Itoa(outputInt)
	case string:
		temp, _ := strconv.Atoi(t)
		outputInt = 7 * temp
		outputString = strconv.Itoa(outputInt)
	default:
		outputInt = 199
		outputString = "99"
	}
	return
}

func computeMathAnyInterface(input interface{}) (output interface{}) {
	switch t := input.(type) {
	case int:
		output = 3 * t
	case string:
		temp, _ := strconv.Atoi(t)
		output = 3 * temp
	default:
		output = 99
	}
	return
}
