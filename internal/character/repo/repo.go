package repo

import (
	"PTOBuilder/internal/character"
	"PTOBuilder/internal/character/model"
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/storage"
	"PTOBuilder/pkg/utils"
	"context"
	"github.com/jackc/pgx/v5"
)

type repo struct {
	log    *logging.Logger
	client storage.Client
}

func NewRepo(log *logging.Logger, client storage.Client) character.Repo {
	return &repo{
		log:    log,
		client: client,
	}
}

func (r *repo) CreateRole(ctx context.Context, role *model.Role) error {
	q := `insert into roles (name) 
		      values($1) 
		  	  returning roles.id;`
	q = utils.FormatQuery(q)
	err := r.client.QueryRow(ctx, q, role.Name).Scan(&role.ID)
	if err != nil {
		r.log.Info(err)
	}
	return err
}

func (r *repo) CreateCharacter(ctx context.Context, character *model.Character) error {
	transaction, err := r.client.Begin(ctx)
	if err != nil {
		r.log.Info(err)
		return err
	}
	q := `insert into characters (name, role_id, image_url, description)
    		  values ($1, $2, $3, $4) 
    		  returning characters.id;`
	q = utils.FormatQuery(q)
	err = transaction.QueryRow(ctx, q, character.Name, character.RoleID, character.ImgUrl, character.Description).Scan(&character.ID)
	if err != nil {
		if err := transaction.Rollback(ctx); err != nil {
			r.log.Info(err)
			return err
		}
		r.log.Info(err)
		return err
	}
	err = r.createStats(ctx, character.Stats, character.ID, transaction)
	if err != nil {
		r.log.Info(err)
		return err
	}
	err = r.createSkills(ctx, character.Skills, character.ID, transaction)
	if err != nil {
		r.log.Info(err)
		return err
	}
	if err := transaction.Commit(ctx); err != nil {
		r.log.Info(err)
		return err
	}
	return nil
}

func (r *repo) createStats(ctx context.Context, stats map[string]*model.Stat, characterID int64, transaction pgx.Tx) error {
	q := `insert into stats (character_id, name, value, scaling)
    		  values ($1, $2, $3, $4) 
    		  returning stats.id;`
	q = utils.FormatQuery(q)
	for key, _ := range stats {
		err := transaction.QueryRow(ctx, q, characterID, key, stats[key].Value, stats[key].Scaling).Scan(&stats[key].ID)
		if err != nil {
			if err := transaction.Rollback(ctx); err != nil {
				r.log.Info(err)
				return err
			}
			r.log.Info(err)
			return err
		}
	}
	return nil
}

func (r *repo) createSkills(ctx context.Context, skill []model.Skill, characterID int64, transaction pgx.Tx) error {
	q := `insert into skills (character_id, name, image_url, description, button)
    		  values ($1, $2, $3, $4, $5)
    		  returning skills.id;`
	q = utils.FormatQuery(q)
	for i, _ := range skill {
		err := transaction.QueryRow(ctx, q, characterID, skill[i].Name, skill[i].ImgUrl, skill[i].Description, skill[i].Button).Scan(&skill[i].ID)
		if err != nil {
			if err := transaction.Rollback(ctx); err != nil {
				r.log.Info(err)
				return err
			}
			r.log.Info(err)
			return err
		}
		err = r.createFormula(ctx, skill[i].Formula, skill[i].ID, transaction)
		if err != nil {
			r.log.Info(err)
			return err
		}
	}
	return nil
}

func (r *repo) createFormula(ctx context.Context, formula []model.Formula, skillID int64, transaction pgx.Tx) error {
	q := `insert into formulas (skill_id, level, formula, stats_name)
    		  values ($1, $2, $3, $4)
    		  returning formulas.id;`
	q = utils.FormatQuery(q)
	for i, _ := range formula {
		err := transaction.QueryRow(ctx, q, skillID, formula[i].Level, formula[i].Formula, formula[i].Stats).Scan(&formula[i].ID)
		if err != nil {
			if err := transaction.Rollback(ctx); err != nil {
				r.log.Info(err)
				return err
			}
			r.log.Info(err)
			return err
		}
	}
	return nil
}

func (r *repo) GetCharacter(ctx context.Context, characterName string) (*model.Character, error) {
	hero := model.Character{}
	q := `select characters.id, roles.id, characters.description, characters.image_url, characters.name
    		  from characters join roles
    		  on roles.id = characters.role_id 
        		  where characters.name = $1;`
	q = utils.FormatQuery(q)
	err := r.client.QueryRow(ctx, q, characterName).Scan(&hero.ID, &hero.RoleID, &hero.Description, &hero.ImgUrl, &hero.Name)
	if err != nil {
		r.log.Info(err)
		return nil, err
	}
	hero.Stats, err = r.getStats(ctx, hero.ID)
	if err != nil {
		r.log.Info(err)
		return nil, err
	}
	return &hero, nil
}

func (r *repo) getStats(ctx context.Context, characterID int64) (map[string]*model.Stat, error) {
	stats := make(map[string]*model.Stat)
	q := `select stats.id, stats.name, stats.scaling, stats.value from stats
    		  where stats.character_id = $1;`
	q = utils.FormatQuery(q)
	rows, err := r.client.Query(ctx, q, characterID)
	if err != nil {
		r.log.Info(err)
		return nil, err
	}
	for rows.Next() {
		var stat model.Stat
		var key string
		err := rows.Scan(&stat.ID, &key, &stat.Scaling, &stat.Value)
		if err != nil {
			r.log.Info(err)
			return nil, err
		}
		stats[key] = &stat
	}
	return stats, nil
}

func (r *repo) getSkills(ctx context.Context, characterID int64) ([]model.Skill, error) {
	skills := make([]model.Skill, 5)
	q := `select skills.id, skills.name, skills.image_url, skills.description, skills.button from skills
    		  where skills.character_id = $1
    		  order by skills.id;`
	q = utils.FormatQuery(q)
	rows, err := r.client.Query(ctx, q, characterID)
	if err != nil {
		r.log.Info(err)
		return nil, err
	}
	for rows.Next() {
		var skill model.Skill
		err := rows.Scan(&skill.ID, skill.Name, skill.ImgUrl, skill.Description, skill.Button)
		if err != nil {
			r.log.Info(err)
			return nil, err
		}
		skill.Formula, err = r.getFormula(ctx, skill.ID)
		if err != nil {
			r.log.Info(err)
			return nil, err
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

func (r *repo) getFormula(ctx context.Context, skillID int64) ([]model.Formula, error) {
	var formulas []model.Formula
	q := `select formulas.id, formulas.level, formulas.formula, formulas.stats_name from formulas
       		  where formulas.skill_id = $1;`
	q = utils.FormatQuery(q)
	rows, err := r.client.Query(ctx, q, skillID)
	if err != nil {
		r.log.Info(err)
		return nil, err
	}
	for rows.Next() {
		var formula model.Formula
		err := rows.Scan(&formula.ID, &formula.Level, &formula.Formula, &formula.Stats)
		if err != nil {
			r.log.Info(err)
			return nil, err
		}
		formulas = append(formulas, formula)
	}
	return formulas, nil
}
