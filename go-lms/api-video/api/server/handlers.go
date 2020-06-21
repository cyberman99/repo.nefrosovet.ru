package server

import (
	"github.com/labstack/echo/v4"

	"repo.nefrosovet.ru/go-lms/api-video/api"
)

// Обработка событий// (POST /api/medooze)
func (s *Server) EventHandler(ctx echo.Context, params api.EventHandlerParams) error {
	return nil
}

// Получение списка вебинаров// (GET /api/v1/webinar)
func (s *Server) WebinarIndex(ctx echo.Context) error {
	return nil
}

// Создание вебинара// (POST /api/v1/webinar)
func (s *Server) WebinarStore(ctx echo.Context) error {
	return nil
}

// Заверешение вебинара// (DELETE /api/v1/webinar/{id})
func (s *Server) WebinarDestroy(ctx echo.Context, id api.Id) error {
	return nil
}

// Получение вебинара// (GET /api/v1/webinar/{id})
func (s *Server) WebinarShow(ctx echo.Context, id api.Id) error {
	return nil
}

// Получение списка мозаик вебинара// (GET /api/v1/webinar/{webinarID}/mosaic/)
func (s *Server) WebinarMosaicIndex(ctx echo.Context, webinarID api.WebinarID) error {
	return nil
}

// Создание мозаики вебинара// (POST /api/v1/webinar/{webinarID}/mosaic/)
func (s *Server) WebinarMosaicStore(ctx echo.Context, webinarID api.WebinarID, params api.WebinarMosaicStoreParams) error {
	return nil
}

// Удаление мозаики вебинара// (DELETE /api/v1/webinar/{webinarID}/mosaic/{mosaicID})
func (s *Server) WebinarMosaicDestroy(ctx echo.Context, webinarID api.WebinarID, mosaicID api.MosaicID) error {
	return nil
}

// Получение мозаики вебинара// (GET /api/v1/webinar/{webinarID}/mosaic/{mosaicID})
func (s *Server) WebinarMosaicShow(ctx echo.Context, webinarID api.WebinarID, mosaicID api.MosaicID) error {
	return nil
}

// Изменение мозаики вебинара// (PATCH /api/v1/webinar/{webinarID}/mosaic/{mosaicID})
func (s *Server) WebinarMosaicUpdate(ctx echo.Context, webinarID api.WebinarID, mosaicID api.MosaicID, params api.WebinarMosaicUpdateParams) error {
	return nil
}

// Получение списка слотов мозаики вебинара// (GET /api/v1/webinar/{webinarID}/mosaic/{mosaicID}/slot)
func (s *Server) WebinarMosaicSlotIndex(ctx echo.Context, webinarID api.WebinarID, mosaicID api.MosaicID) error {
	return nil
}

// Изменение слота мозаики вебинара// (PATCH /api/v1/webinar/{webinarID}/mosaic/{mosaicID}/slot/{slotID})
func (s *Server) WebinarMosaicSlotUpdate(ctx echo.Context, webinarID api.WebinarID, mosaicID api.MosaicID, slotID api.SlotID, params api.WebinarMosaicSlotUpdateParams) error {
	return nil
}

// Получение списка юзеров мозаики вебинара// (GET /api/v1/webinar/{webinarID}/mosaic/{mosaicID}/user)
func (s *Server) WebinarMosaicUserIndex(ctx echo.Context, webinarID api.WebinarID, mosaicID api.MosaicID) error {
	return nil
}

// Получение списка мозаик вебинара// (POST /api/v1/webinar/{webinarID}/mosaic/{mosaicID}/user)
func (s *Server) WebinarMosaicUserStore(ctx echo.Context, webinarID api.WebinarID, mosaicID api.MosaicID, params api.WebinarMosaicUserStoreParams) error {
	return nil
}

// Получение юзера мозаики вебинара// (DELETE /api/v1/webinar/{webinarID}/mosaic/{mosaicID}/user/{userID})
func (s *Server) WebinarMosaicUserShow(ctx echo.Context, webinarID api.WebinarID, mosaicID api.MosaicID, userID api.UserID) error {
	return nil
}

// Получение списка потоков// (GET /api/v1/webinar/{webinarID}/stream)
func (s *Server) WebinarStreamIndex(ctx echo.Context, webinarID api.WebinarID) error {
	return nil
}

// Создание потока// (POST /api/v1/webinar/{webinarID}/stream)
func (s *Server) WebinarStreamStore(ctx echo.Context, webinarID api.WebinarID, params api.WebinarStreamStoreParams) error {
	return nil
}

// Удаление потока// (DELETE /api/v1/webinar/{webinarID}/stream/{streamID})
func (s *Server) WebinarStreamDestroy(ctx echo.Context, webinarID api.WebinarID, streamID api.StreamID) error {
	return nil
}

// Изменение данных потока// (PATCH /api/v1/webinar/{webinarID}/stream/{streamID})
func (s *Server) WebinarStreamUpdate(ctx echo.Context, webinarID api.WebinarID, streamID api.StreamID, params api.WebinarStreamUpdateParams) error {
	return nil
}

// Создание мозаики юзера вебинара// (POST /api/v1/webinar/{webinarID}/user/{userID}/mosaic)
func (s *Server) WebinarUserMosaicStore(ctx echo.Context, webinarID api.WebinarID, userID api.UserID, params api.WebinarUserMosaicStoreParams) error {
	return nil
}
