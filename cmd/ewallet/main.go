package main

import (
	"ewallet/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	//envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting ewallet")

	//TODO: init storage: postgresql psx

	//TODO: init router: http-server

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}

// обработка транзакций платёжной системы. В виде HTTP сервера на REST API
// с 4 методами:

/*TODO: запрос создание кошелька
	POST /api/v1/wallet
параметры запроса - нет
ответ JSON с сотсоянием созданного кошелька:
	-- id - строковый id кошелька, генерируется сервером
	-- balance - дробное, баланс кошелька
*созданный кошелёк должен иметь 100уе на балансе
*/

/*TODO: запрос перевод с кошелька на кошелек
	POST /api/v1/wallet/{walletId}/send
параметры запроса:
	-- walletId - строковый id указан в пути запроса
	-- JSON-объект в теле запроса с параметрами:
		to - ID кошелька, куда переводим
		amount - сумма перевода
ответ:
	-- 200 успешно
	-- 404 исходящий кошель не найден
	-- 400 целевой не найден, не хватает баланса на исходящем
*/

/*TODO: запрос получение истории входящих и исходящих транзакций
	GET /api/v1/wallet/{walletId}/history
параметры запроса:
	walletId - строковый ID кошелька указан в пути запроса
ответ:
	-- 200 если кошель найден
		ответ содержит массив JSON-объектов с вход и исход транзакциями кошеля
		каждый объект содержит параметры:
			time - RFC 3339 дата и время перевода
			from - ID исходящего
			to   - ID входящего
			amount - сумма перевода. Дробное
	-- 404 указанный кошель не найден
*/

/*TODO: запрос получение текущего состояния кошелька
	GET /api/v1/{walletId}
параметры запроса:
	walletId - строковый id кошелька, указан в пути запроса
ответ:
	-- 200 найден кошель
		ответ содержит в теле JSON-объект с текущим состоянием кошеля
		содержит параметры:
			id - строковый ID кошелька. Генерируется сервером
			balance - дробное число, баланс кошелька
	-- 404 кошель не найден
*/

/*
	безопасность - 	не должно быть уязвимостей, чтобы можно было произвольно
					менять данные в базе
	персистентность - 	данные и изменения не должны теряться при перезапуске
						приложения

	плюсом будет - наличие в решении dockerfile для сборки контейнера с прилож.
				 - хранение в Git
*/
