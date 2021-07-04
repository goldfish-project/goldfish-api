package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"goldfish-api/internal/core/domain"
)

type WorkspaceRepository struct {
	db *pg.DB
}

// NewWorkspaceRepository creates a new WorkspaceRepository an returns it
func NewWorkspaceRepository(db *pg.DB) WorkspaceRepository {
	return WorkspaceRepository{
		db: db,
	}
}

// Get retrieves a workspace by its id
func (repo *WorkspaceRepository) Get(workspaceId uuid.UUID) (domain.Workspace, error) {
	workspace := &domain.Workspace{WorkspaceId: workspaceId}

	// select workspace from database by given id
	if err := repo.db.Model(workspace).WherePK().Select(); err != nil {
		return domain.Workspace{}, nil
	}

	return *workspace, nil
}

// Save creates a new workspace
func (repo *WorkspaceRepository) Save(workspace *domain.Workspace, ownerUserId string) error {
	// parse into database model
	workspaceDB := Workspace{}
	workspaceDB.from(workspace, workspace.Owner.UserId)

	// save given instance of workspace into database
	_, err := repo.db.Model(&workspaceDB).Insert();

	if err != nil {
		return err
	}

	//TODO: save id to parsed workspace paramater

	return nil
}

// Delete deletes a given workspace
func (repo *WorkspaceRepository) Delete(workspaceId uuid.UUID) error {
	workspaceDB := Workspace{
		WorkspaceId: workspaceId,
	}

	// delete database entry
	if _, err := repo.db.Model(workspaceDB).ForceDelete(); err != nil {
		return err
	}

	return nil
}

// GetForUser retrieves all workspaces for a given user no matter whether their the owner or an collaborator
func (repo *WorkspaceRepository) GetForUser(userId string) ([]domain.Workspace, error) {
	var ownWorkspaces []Workspace
	var collaboratorWorkspaces []Workspace

	// select all workspaces which are owner by the given user from database
	if err := repo.db.Model(&ownWorkspaces).Where("? = ?", pg.Ident("owner_id"), userId).Select(); err != nil {
		return []domain.Workspace{}, err
	}

	// TODO: select all workspaces where the given user is a collaborator

	return []domain.Workspace{}, nil
}