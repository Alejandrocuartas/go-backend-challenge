package repository

import (
	"fmt"
	"go-backend-challenge/internal/model"

	utils "github.com/Alejandrocuartas/core-utils-private-library"

	"github.com/Alejandrocuartas/core-models-private-library/models/agencies"
	"github.com/Alejandrocuartas/core-models-private-library/models/campaign_creator_social_network_actions"
	"github.com/Alejandrocuartas/core-models-private-library/models/campaigns"
	"github.com/Alejandrocuartas/core-models-private-library/models/companies"
	"github.com/Alejandrocuartas/core-models-private-library/models/creator_social_networks"
	uar "github.com/Alejandrocuartas/core-models-private-library/models/user_agency_relations"
	"github.com/Alejandrocuartas/core-models-private-library/models/users"

	"github.com/jinzhu/gorm"
)

type AgenciesDbRepository struct {
	*gorm.DB
}

func (c AgenciesDbRepository) GetUserById(
	userId uint,
) (
	user users.User,
	err error,
) {
	err = c.Table("users").
		Where("id = ?", userId).
		First(&user).Error

	return user, err
}

func (c AgenciesDbRepository) GetCampaignById(
	id uint,
) (
	campaigns.Campaign,
	error,
) {
	var campaign campaigns.Campaign

	err := c.Table("campaigns").
		Where("id = ?", id).
		First(&campaign).Error
	return campaign, err
}

func (c AgenciesDbRepository) GetCompanyById(
	id uint,
) (
	companies.Company,
	error,
) {
	var company companies.Company

	err := c.Table("companies").
		Where("id = ?", id).
		First(&company).Error
	return company, err
}

func (c AgenciesDbRepository) GetSocialNetworkById(
	id uint,
) (
	creator_social_networks.CreatorSocialNetworkAccount,
	error,
) {
	var creatorSocialNetworkAccount creator_social_networks.CreatorSocialNetworkAccount

	err := c.Table("creator_social_networks").
		Where("id = ?", id).
		First(&creatorSocialNetworkAccount).Error
	return creatorSocialNetworkAccount, err
}

func (c AgenciesDbRepository) CreateCampaign(
	u campaigns.Campaign,
) (
	campaigns.Campaign,
	error,
) {
	c.Exec(
		fmt.Sprintf(
			"SELECT setval('%s_id_seq', (select max(id) from %s) + 1, FALSE);",
			"campaigns",
			"campaigns",
		),
	)
	if err := c.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (c AgenciesDbRepository) CreateAction(
	u campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions,
) (
	campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions,
	error,
) {
	c.Exec(
		fmt.Sprintf(
			"SELECT setval('%s_id_seq', (select max(id) from %s) + 1, FALSE);",
			"campaign_creator_social_network_actions",
			"campaign_creator_social_network_actions",
		),
	)
	if err := c.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (c AgenciesDbRepository) CreateAgency(
	u agencies.Agency,
) (
	agencies.Agency,
	error,
) {
	c.Exec(
		fmt.Sprintf(
			"SELECT setval('%s_id_seq', (select max(id) from %s) + 1, FALSE);",
			"agencies",
			"agencies",
		),
	)
	if err := c.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (c AgenciesDbRepository) CreateUserAgencyRelation(
	u uar.UserAgencyRelation,
) (
	uar.UserAgencyRelation,
	error,
) {
	c.Exec(
		fmt.Sprintf(
			"SELECT setval('%s_id_seq', (select max(id) from %s) + 1, FALSE);",
			"user_agency_relations",
			"user_agency_relations",
		),
	)
	if err := c.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (c AgenciesDbRepository) CreateUser(
	u users.User,
) (
	users.User,
	error,
) {
	c.Exec(
		fmt.Sprintf(
			"SELECT setval('%s_id_seq', (select max(id) from %s) + 1, FALSE);",
			"users",
			"users",
		),
	)
	if err := c.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (c AgenciesDbRepository) CreateCampaignCreatorSocialNetworkAction(
	u campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions,
) (
	campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions,
	error,
) {
	c.Exec(
		fmt.Sprintf(
			"SELECT setval('%s_id_seq', (select max(id) from %s) + 1, FALSE);",
			"campaign_creator_social_network_actions",
			"campaign_creator_social_network_actions",
		),
	)
	if err := c.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (a AgenciesDbRepository) ListActionsCampaign(campaignId uint) (
	[]campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions,
	error,
) {

	var campaignActions []campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions

	selectColumns := `
	campaign_creator_social_network_actions.code_name as action_name,
	campaign_creator_social_network_actions.quantity as action_quantity,
	campaign_creator_social_network_actions.cost_price as action_price,
	campaign_creator_social_network_actions.cost_currency as action_cost_currency
`

	mtdt := a.Table("campaign_creator_social_network_actions").
		Select(selectColumns).
		Where("campaign_creator_social_network_actions.campaign_id = ?", campaignId).
		Where("campaign_creator_social_network_actions.deleted_at is null")

	if mtdt.Error != nil {
		return nil, mtdt.Error
	}

	rows, err := mtdt.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var action campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions
		err := rows.Scan(
			&action.CodeName,
			&action.Quantity,
			&action.CostPrice,
			&action.CostCurrency,
		)
		if err != nil {
			return nil, err
		}
		campaignActions = append(campaignActions, action)
	}

	return campaignActions, nil

}

func (a AgenciesDbRepository) ListAgencyCampaigns(
	pagination utils.GormPaginationData,
	agencyId uint,
	keywordSearch string,
) (
	utils.GormPaginationData,
	error,
) {

	if pagination.Sort == "" {
		pagination.Sort = "id desc"
	}

	var results []model.ListCampaignsResponseModel

	selectColumns := `
		campaigns.id as campaign_id,
		campaigns.created_at as campaign_created_at,
		campaigns.updated_at as campaign_updated_at,
		campaigns.deleted_at as campaign_deleted_at,
		campaigns.name as campaign_name,
		campaigns.initial_date as campaign_initial_date,
		campaigns.final_date as campaign_final_date,
		campaigns.budget as campaign_budget,
		campaigns.currency as campaign_currency,
		campaigns.agency_id as campaign_agency_id,
		campaigns.manager_id as campaign_manager_id,
		campaigns.company_id as campaign_company_id,
		campaigns.bundle_id as campaign_bundle_id,
		users.first_name as manager_name,
		users.email as manager_email,
		companies.name as company_name,
		companies.email as company_email
	`

	tx := a.Table("campaigns").
		Select(selectColumns).
		Joins("join users on users.id = campaigns.manager_id").
		Joins("join companies on companies.id = campaigns.company_id").
		Where("campaigns.agency_id = ?", agencyId).
		Where("campaigns.deleted_at is null")

	if keywordSearch != "" {
		tx = tx.
			Where(
				`campaigns.name ilike ? 
			or users.first_name ilike ?
			or companies.name ilike ?`,
				"%"+keywordSearch+"%",
				"%"+keywordSearch+"%",
				"%"+keywordSearch+"%",
			)
	}

	pagination.Sort = "campaigns.id desc"

	tx = tx.Scopes(utils.Paginate(results, &pagination, tx)).
		Scan(&results)
	pagination.Rows = results

	if pagination.Rows == nil {
		pagination.Rows = make([]interface{}, 0)
		return utils.GormPaginationData{}, tx.Error
	}
	for i := 0; i < len(results); i++ {

		acts, e := a.ListActionsCampaign(results[i].CampaignID)
		if e != nil {
			return utils.GormPaginationData{}, e
		}
		for _, act := range acts {
			if act.CodeName != "" {
				results[i].CampaignCreatorSocialNetworkActions = acts
			}
		}
	}

	pagination.Rows = results

	return pagination, tx.Error

}

func (c AgenciesDbRepository) IsAnManagerAgencyRelation(
	agencyId uint,
	managerId uint,
) (
	bool,
	error,
) {

	var relationExists bool
	var res uar.UserAgencyRelation

	err := c.DB.
		Table("user_agency_relations").
		Where("user_id = ?", managerId).
		Where("agency_id = ?", agencyId).
		First(&res).Error

	if res.UserId != 0 {
		relationExists = true
	}

	return relationExists, err
}

func (c AgenciesDbRepository) IsAnAgencyCompanyRelation(
	agencyId uint,
	companyId uint,
) (
	bool,
	error,
) {

	var relationExists bool
	var company companies.Company

	err := c.DB.
		Table(company.TableName()).
		Where("id = ?", companyId).
		Where("agency_id = ?", agencyId).
		First(&company).Error

	if company.ID != 0 {
		relationExists = true
	}

	return relationExists, err
}

func (c AgenciesDbRepository) UpdateCampaign(
	u campaigns.Campaign,
) (
	campaigns.Campaign,
	error,
) {
	err := c.Save(&u).Error
	return u, err
}
