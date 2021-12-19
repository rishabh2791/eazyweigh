package application

import "eazyweigh/infrastructure/persistance"

type AppStore struct {
	repoStore *persistance.RepoStore
}

func NewAppStore(repoStore *persistance.RepoStore) *AppStore {
	appStore := AppStore{}
	appStore.repoStore = repoStore
	return &appStore
}
