package main

import (
	"context"
	"ewallet/internal/config"
	"ewallet/internal/logger"
	operationDb "ewallet/internal/models/operation/db"
	"ewallet/internal/storage/postgresql"
	"time"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	log.Info("starting ewallet")

	pgxPoolClient, err := postgresql.NewClient(context.TODO(), cfg)
	if err != nil {
		log.Error("error with getting pgx pool: ", err)
	}

	//repositoryWallet := walletDb.NewRepository(pgxPoolClient, log)
	repositoryOperation := operationDb.NewRepository(pgxPoolClient, log)

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Error("err with get location", err)
	}
	timeToInsert := time.Date(2024, time.July, 15, 10, 30, 0, 0, loc)
	isCreated, err := repositoryOperation.Create(context.TODO(), timeToInsert,
		"752224ff-bfae-4581-8bce-3ab9911e7cd2", "6bed2193-1359-4e46-b6dd-490194a717b8", 40)
	if err != nil {
		log.Error("err with get location", err)
	}
	if isCreated {
		log.Info("operation creating successful")
	}
	//TODO: init router: http-server

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
