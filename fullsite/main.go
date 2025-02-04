package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var indexTemplate *template.Template

func init() {
	indexTemplate = template.Must(template.ParseFiles(filepath.Join("template", "index.html")))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // Устанавливаем кодировку UTF-8

	switch r.Method {
	case "GET":
		// Рендерим шаблон
		err := indexTemplate.Execute(w, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при выполнении шаблона: %v", err), http.StatusInternalServerError)
			return
		}
	case "POST":
		// Разбираем форму
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при разборе формы: %v", err), http.StatusBadRequest)
			return
		}

		name := r.FormValue("username")
		email := r.FormValue("useremail")
		message := r.FormValue("usermessage")

		log.Printf("%s %s %s", name, email, message)

		// Вызываем функцию отправки письма
		sendEmailCallback(name, email, message)

		// Выводим сообщение об успешной отправке
		w.Write([]byte(`
    <!DOCTYPE html>
    <html lang="ru">
    <head>
        <meta charset="UTF-8">
        <title>Успешная отправка</title>
        <script>
            setTimeout(function(){
                window.location.href = "/";
            }, 3000); // Переход на главную страницу через 3 секунды
        </script>
    </head>
    <body>
        <h1>Ваше письмо было успешно отправлено!</h1>
        <p>Вы будете перенаправлены на главную страницу через 3 секунды.</p>
    </body>
    </html>
    `))
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func AppointmentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(filepath.Join("template", "form.html"))
	if err != nil {
		fmt.Fprintf(w, "Ошибка при парсинге шаблона: %v", err)
		return
	}
	// Выполняем рендеринг шаблона
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Ошибка при выполнении шаблона: %v", err)
		return
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	username := r.Form.Get("username")
	surname := r.Form.Get("surname")
	useremail := r.Form.Get("useremail")
	userservice := r.Form.Get("userservice")
	date := r.Form.Get("date")
	time := r.Form.Get("time")
	usermessage := r.Form.Get("usermessage")

	log.Printf("%s %s %s %s %s %s %s", username, surname, useremail, userservice,
		date, time, usermessage)

	// Формируем полное имя пользователя
	fullName := username + " " + surname

	// Вызываем функцию отправки письма
	err = sendEmailAppointment(fullName, useremail, userservice, date, time, usermessage)
	if err != nil {
		log.Println(err)
		return
	}
	// Выводим сообщение об успешной отправке
	fmt.Fprint(w, `
    <!DOCTYPE html>
    <html lang="ru">
    <head>
        <meta charset="UTF-8">
        <title>Успешная отправка</title>
        <script>
            setTimeout(function(){
                window.location.href = "/";
            }, 3000); // Переход на главную страницу через 3 секунды
        </script>
    </head>
    <body>
        <h1>Ваше письмо было успешно отправлено!</h1>
        <p>Вы будете перенаправлены на главную страницу через 3 секунды.</p>
    </body>
    </html>
    `)
}

func main() {
	// Создаем файл для логирования
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла для логирования: %v", err)
	}
	defer logFile.Close()

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/appointment", AppointmentHandler)
	http.HandleFunc("/submit", SubmitHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Запуск сервера на http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	// Запись логов
	log.Println("Это тестовый лог")
	log.Printf("Логирование значения: %d\n", 42)

}
