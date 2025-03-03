package domain

//
//type Domain interface {
//	CreateSubscriptionType(ctx context.Context, sub models.SubscriptionType) error
//	CreateSubscriptionPlan(ctx context.Context, plan models.SubscriptionPlan) error
//
//	DeleteSubscriptionType(ctx context.Context, id uuid.UUID) error
//	DeleteSubscriptionPlan(ctx context.Context, id uuid.UUID) error
//
//	UpdateSubscriptionType() error
//	UpdateSubscriptionPlan() error
//}
//
//type domain struct {
//	Infra *infra.Infra
//	log   *logrus.Logger
//}
//
//func NewDomain(infra *infra.Infra, log *logrus.Logger) (Domain, error) {
//	return &domain{
//		Infra: infra,
//		log:   log,
//	}, nil
//}
//
//func (d *domain) CreateSubscriptionType(ctx context.Context, subType models.SubscriptionType) error {
//	return d.Infra.Data.SubTypes.Create(ctx, subType)
//}
//
//func (d *domain) CreateSubscriptionPlan(ctx context.Context, plan models.SubscriptionPlan) error {
//	return d.Infra.Data.SubPlans.Create(ctx, plan)
//}
//
//func (d *domain) UpdateSubscriptionType(ctx context.Context, id uuid.UUID) error {
//	return d.Infra.Data.SubTypes.Filter(map[string]any{
//		"id": id.String(),
//	}).Update(ctx
//}
//
//func (d *domain) UpdateSubscriptionPlan(ctx context.Context, id uuid.UUID) error {
//	return d.Infra.Data.SubPlans.Filter(map[string]any{
//		"id": id.String(),
//	}).DeleteByID()
