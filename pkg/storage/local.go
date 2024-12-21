package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"kotakemail.id/config"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/logger"
)

type LocalStorage struct {
	name   string
	cfg    *config.StorageConfig
	logger *logger.Logger
}

func NewLocalStorage(cfg *config.StorageConfig, appLogger *logger.Logger) Storage {
	return &LocalStorage{
		name:   cfg.Name,
		cfg:    cfg,
		logger: appLogger,
	}
}

func (s *LocalStorage) Name() string {
	return s.name
}

func (s *LocalStorage) Write(ctx *appcontext.AppContext, path string, options *WriterOptions) (io.WriteCloser, error) {
	path = s.fullPath(path)
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			s.logger.Error().Err(err).Msg("failed to create directory")
			return nil, err
		}
	}

	f, err := os.Create(path)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to create file")
		return nil, err
	}

	if options != nil {
		if !options.Attributes.CreationTime.IsZero() {
			options.Attributes.ModTime = options.Attributes.CreationTime
		}
		if options.Attributes.ModTime.IsZero() {
			if err := os.Chtimes(path, options.Attributes.ModTime, options.Attributes.ModTime); err != nil {
				s.logger.Error().Err(err).Msg("failed to set file mod time")
				return nil, err
			}
		}
	}

	return f, nil
}

func (s *LocalStorage) Read(ctx *appcontext.AppContext, path string, options *ReaderOptions) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, s.wrapError(path, err)
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, s.wrapError(path, err)
	}

	return &File{
		ReadCloser: f,
		Attributes: Attributes{
			ModTime: stat.ModTime(),
			Size:    stat.Size(),
		},
	}, nil
}

func (s *LocalStorage) Delete(ctx *appcontext.AppContext, path string) error {
	path = s.fullPath(path)
	if err := os.Remove(path); err != nil {
		return s.wrapError(path, err)
	}

	return nil
}

func (s *LocalStorage) GetURL(ctx *appcontext.AppContext, path string) (string, error) {
	return fmt.Sprintf("%s/%s", s.cfg.DeliveryBasePath, path), nil
}

func (s *LocalStorage) fullPath(path string) string {
	return fmt.Sprintf("%s/%s", s.cfg.BasePath, path)
}

func (s *LocalStorage) wrapError(path string, err error) error {
	if os.IsNotExist(err) {
		err = &notExistError{
			Path: path,
		}
	}

	s.logger.Error().Err(err).Msg("failed to perform operation")
	return err
}
