package res_partner

import "easonsu/supo/db"
import (
	"time"
	"easonsu/supo/z-core/base"
)

type ResPartner struct {
	model.BaseModel
	Name        string    `json:"name" gorm:"column:name;type:varchar;"`
	DisplayName string    `json:"display_name" gorm:"column:display_name;type:varchar;"` // compute field display_name store is True
	Date        time.Time `json:"date" gorm:"column:date;type:timestamp(0) with time zone;"`
	//Title   *ResPartnerTitle `json:"title" gorm:"column:title;type:int;"`
	ParentId *ResPartner   `json:"parent_id" gorm:"column:parent_id;type:int;"`
	ChildIds *[]ResPartner `json:"child_ids" gorm:"column:child_ids;type:;"`
	Ref      string        `json:"ref" gorm:"column:ref;type:varchar;"`
	Lang     string        `json:"lang" gorm:"column:lang;type:varchar;"`
	Tz       string        `json:"tz" gorm:"column:tz;type:varchar;"`
	TzOffset string        `json:"tz_offset" gorm:"column:tz_offset;type:varchar;"` // compute field tz_offset store is False
	//UserId   *ResUsers `json:"user_id" gorm:"column:user_id;type:int;"`
	Vat     string            `json:"vat" gorm:"column:vat;type:varchar;"`
	BankIds *[]ResPartnerBank `json:"bank_ids" gorm:"column:bank_ids;type:;"`
	Website string            `json:"website" gorm:"column:website;type:varchar;"`
	Comment string            `json:"comment" gorm:"column:comment;type:varchar;"`
	//CategoryId   *[]ResPartnerCategory `json:"category_id" gorm:"column:category_id;type:;"`
	CreditLimit float32 `json:"credit_limit" gorm:"column:credit_limit;type:;"`
	Barcode     string  `json:"barcode" gorm:"column:barcode;type:varchar;"`
	Active      bool    `json:"active" gorm:"column:active;type:bool;"`
	Customer    bool    `json:"customer" gorm:"column:customer;type:bool;"`
	Supplier    bool    `json:"supplier" gorm:"column:supplier;type:bool;"`
	Employee    bool    `json:"employee" gorm:"column:employee;type:bool;"`
	Function    string  `json:"function" gorm:"column:function;type:varchar;"`
	Type        string  `json:"type" gorm:"column:type;type:varchar;"`
	Street      string  `json:"street" gorm:"column:street;type:varchar;"`
	Street2     string  `json:"street2" gorm:"column:street2;type:varchar;"`
	Zip         string  `json:"zip" gorm:"column:zip;type:varchar;"`
	//City   *ResCountryStateCity `json:"city" gorm:"column:city;type:int;"`
	//StateId   *ResCountryState `json:"state_id" gorm:"column:state_id;type:int;"`
	//CountryId   *ResCountry `json:"country_id" gorm:"column:country_id;type:int;"`
	Email          string `json:"email" gorm:"column:email;type:varchar;"`
	EmailFormatted string `json:"email_formatted" gorm:"column:email_formatted;type:varchar;"` // compute field email_formatted store is False
	Phone          string `json:"phone" gorm:"column:phone;type:varchar;"`
	Mobile         string `json:"mobile" gorm:"column:mobile;type:varchar;"`
	IsCompany      bool   `json:"is_company" gorm:"column:is_company;type:bool;"`
	//IndustryId   *ResPartnerIndustry `json:"industry_id" gorm:"column:industry_id;type:int;"`
	//CompanyType   None `json:"company_type" gorm:"column:company_type;type:;"`  // compute field company_type store is False
	//CompanyId   *ResCompany `json:"company_id" gorm:"column:company_id;type:int;"`
	//Color   int `json:"color" gorm:"column:color;type:int;"`
	//UserIds   *[]ResUsers `json:"user_ids" gorm:"column:user_ids;type:;"`
	PartnerShare          bool        `json:"partner_share" gorm:"column:partner_share;type:bool;"`                        // compute field partner_share store is True
	ContactAddress        string      `json:"contact_address" gorm:"column:contact_address;type:varchar;"`                 // compute field contact_address store is False
	CommercialPartnerId   *ResPartner `json:"commercial_partner_id" gorm:"column:commercial_partner_id;type:int;"`         // compute field commercial_partner_id store is True
	CommercialCompanyName string      `json:"commercial_company_name" gorm:"column:commercial_company_name;type:varchar;"` // compute field commercial_company_name store is True
	CompanyName           string      `json:"company_name" gorm:"column:company_name;type:varchar;"`
	Image                 string      `json:"image" gorm:"column:image;type:varchar;"`
	ImageMedium           string      `json:"image_medium" gorm:"column:image_medium;type:varchar;"`
	ImageSmall            string      `json:"image_small" gorm:"column:image_small;type:varchar;"`
	Self                  *ResPartner `json:"self" gorm:"column:self;type:int;"` // compute field self store is False
	Id                    int         `json:"id" gorm:"column:id;type:int;;primary_key;AUTO_INCREMENT"`
	LastUpdate            time.Time   `json:"__last_update" gorm:"column:__last_update;type:timestamp(0) with time zone;"` // compute field __last_update store is False
	ImStatus              string      `json:"im_status" gorm:"column:im_status;type:varchar;"`                             // compute field im_status store is False
	//ActivityIds   *[]MailActivity `json:"activity_ids" gorm:"column:activity_ids;type:;"`
	//ActivityState   None `json:"activity_state" gorm:"column:activity_state;type:;"`  // compute field activity_state store is False
	MessageIsFollower bool `json:"message_is_follower" gorm:"column:message_is_follower;type:bool;"` // compute field message_is_follower store is False
	//MessageFollowerIds   *[]MailFollowers `json:"message_follower_ids" gorm:"column:message_follower_ids;type:;"`
	//MessagePartnerIds   *[]ResPartner `json:"message_partner_ids" gorm:"column:message_partner_ids;type:;"`  // compute field message_partner_ids store is False
	//MessageChannelIds   *[]MailChannel `json:"message_channel_ids" gorm:"column:message_channel_ids;type:;"`  // compute field message_channel_ids store is False
	//MessageIds   *[]MailMessage `json:"message_ids" gorm:"column:message_ids;type:;"`
	MessageLastPost          time.Time `json:"message_last_post" gorm:"column:message_last_post;type:timestamp(0) with time zone;"`
	MessageUnread            bool      `json:"message_unread" gorm:"column:message_unread;type:bool;"`                        // compute field message_unread store is False
	MessageUnreadCounter     int       `json:"message_unread_counter" gorm:"column:message_unread_counter;type:int;"`         // compute field message_unread_counter store is False
	MessageNeedaction        bool      `json:"message_needaction" gorm:"column:message_needaction;type:bool;"`                // compute field message_needaction store is False
	MessageNeedactionCounter int       `json:"message_needaction_counter" gorm:"column:message_needaction_counter;type:int;"` // compute field message_needaction_counter store is False
	MessageBounce            int       `json:"message_bounce" gorm:"column:message_bounce;type:int;"`
	OptOut                   bool      `json:"opt_out" gorm:"column:opt_out;type:bool;"`
	//ChannelIds   *[]MailChannel `json:"channel_ids" gorm:"column:channel_ids;type:;"`
	SignupToken          string    `json:"signup_token" gorm:"column:signup_token;type:varchar;"`
	SignupType           string    `json:"signup_type" gorm:"column:signup_type;type:varchar;"`
	SignupExpiration     time.Time `json:"signup_expiration" gorm:"column:signup_expiration;type:timestamp(0) with time zone;"`
	SignupValid          bool      `json:"signup_valid" gorm:"column:signup_valid;type:bool;"` // compute field signup_valid store is False
	SignupUrl            string    `json:"signup_url" gorm:"column:signup_url;type:varchar;"`  // compute field signup_url store is False
	CalendarLastNotifAck time.Time `json:"calendar_last_notif_ack" gorm:"column:calendar_last_notif_ack;type:timestamp(0) with time zone;"`
	//WebsiteMessageIds   *[]MailMessage `json:"website_message_ids" gorm:"column:website_message_ids;type:;"`
	//PropertyProductPricelist   *ProductPricelist `json:"property_product_pricelist" gorm:"column:property_product_pricelist;type:int;"`  // compute field property_product_pricelist store is False
	//TeamId   *CrmTeam `json:"team_id" gorm:"column:team_id;type:int;"`
	Credit        float32 `json:"credit" gorm:"column:credit;type:;"` // compute field credit store is False
	Debit         float32 `json:"debit" gorm:"column:debit;type:;"`   // compute field debit store is False
	DebitLimit    float32 `json:"debit_limit" gorm:"column:debit_limit;type:;"`
	TotalInvoiced float32 `json:"total_invoiced" gorm:"column:total_invoiced;type:;"` // compute field total_invoiced store is False
	//CurrencyId   *ResCurrency `json:"currency_id" gorm:"column:currency_id;type:int;"`  // compute field currency_id store is False
	ContractsCount   int `json:"contracts_count" gorm:"column:contracts_count;type:int;"`       // compute field contracts_count store is False
	JournalItemCount int `json:"journal_item_count" gorm:"column:journal_item_count;type:int;"` // compute field journal_item_count store is False
	//PropertyAccountPayableId   *AccountAccount `json:"property_account_payable_id" gorm:"column:property_account_payable_id;type:int;not null"`  // compute field property_account_payable_id store is False
	//PropertyAccountReceivableId   *AccountAccount `json:"property_account_receivable_id" gorm:"column:property_account_receivable_id;type:int;not null"`  // compute field property_account_receivable_id store is False
	//PropertyAccountPositionId   *AccountFiscalPosition `json:"property_account_position_id" gorm:"column:property_account_position_id;type:int;"`  // compute field property_account_position_id store is False
	//PropertyPaymentTermId   *AccountPaymentTerm `json:"property_payment_term_id" gorm:"column:property_payment_term_id;type:int;"`  // compute field property_payment_term_id store is False
	//PropertySupplierPaymentTermId   *AccountPaymentTerm `json:"property_supplier_payment_term_id" gorm:"column:property_supplier_payment_term_id;type:int;"`  // compute field property_supplier_payment_term_id store is False
	//RefCompanyIds   *[]ResCompany `json:"ref_company_ids" gorm:"column:ref_company_ids;type:;"`
	//HasUnreconciledEntries   bool `json:"has_unreconciled_entries" gorm:"column:has_unreconciled_entries;type:bool;"`  // compute field has_unreconciled_entries store is False
	//LastTimeEntriesChecked   time.Time `json:"last_time_entries_checked" gorm:"column:last_time_entries_checked;type:timestamp(0) with time zone;"`
	//InvoiceIds   *[]AccountInvoice `json:"invoice_ids" gorm:"column:invoice_ids;type:;"`
	//ContractIds   *[]AccountAnalyticAccount `json:"contract_ids" gorm:"column:contract_ids;type:;"`
	BankAccountCount int `json:"bank_account_count" gorm:"column:bank_account_count;type:int;"` // compute field bank_account_count store is False
	//Trust   None `json:"trust" gorm:"column:trust;type:;"`  // compute field trust store is False
	InvoiceWarn    string `json:"invoice_warn" gorm:"column:invoice_warn;type:varchar;not null"`
	InvoiceWarnMsg string `json:"invoice_warn_msg" gorm:"column:invoice_warn_msg;type:varchar;"`
	//TaskIds   *[]ProjectTask `json:"task_ids" gorm:"column:task_ids;type:;"`
	TaskCount int `json:"task_count" gorm:"column:task_count;type:int;"` // compute field task_count store is False
	//PropertyStockCustomer   *StockLocation `json:"property_stock_customer" gorm:"column:property_stock_customer;type:int;"`  // compute field property_stock_customer store is False
	//PropertyStockSupplier   *StockLocation `json:"property_stock_supplier" gorm:"column:property_stock_supplier;type:int;"`  // compute field property_stock_supplier store is False
	PickingWarn    string `json:"picking_warn" gorm:"column:picking_warn;type:varchar;not null"`
	PickingWarnMsg string `json:"picking_warn_msg" gorm:"column:picking_warn_msg;type:varchar;"`
	//PaymentTokenIds   *[]PaymentToken `json:"payment_token_ids" gorm:"column:payment_token_ids;type:;"`
	PaymentTokenCount int `json:"payment_token_count" gorm:"column:payment_token_count;type:int;"` // compute field payment_token_count store is False
	SaleOrderCount    int `json:"sale_order_count" gorm:"column:sale_order_count;type:int;"`       // compute field sale_order_count store is False
	//SaleOrderIds   *[]SaleOrder `json:"sale_order_ids" gorm:"column:sale_order_ids;type:;"`
	SaleWarn    string `json:"sale_warn" gorm:"column:sale_warn;type:varchar;not null"`
	SaleWarnMsg string `json:"sale_warn_msg" gorm:"column:sale_warn_msg;type:varchar;"`
	//PropertyPurchaseCurrencyId   *ResCurrency `json:"property_purchase_currency_id" gorm:"column:property_purchase_currency_id;type:int;"`  // compute field property_purchase_currency_id store is False
	PurchaseOrderCount   int    `json:"purchase_order_count" gorm:"column:purchase_order_count;type:int;"`     // compute field purchase_order_count store is False
	SupplierInvoiceCount int    `json:"supplier_invoice_count" gorm:"column:supplier_invoice_count;type:int;"` // compute field supplier_invoice_count store is False
	PurchaseWarn         string `json:"purchase_warn" gorm:"column:purchase_warn;type:varchar;not null"`
	PurchaseWarnMsg      string `json:"purchase_warn_msg" gorm:"column:purchase_warn_msg;type:varchar;"`
	//PropertyDeliveryCarrierId   *DeliveryCarrier `json:"property_delivery_carrier_id" gorm:"column:property_delivery_carrier_id;type:int;"`  // compute field property_delivery_carrier_id store is False
	Openid          string `json:"openid" gorm:"column:openid;type:varchar;"`
	DisplayNameChar string `json:"display_name_char" gorm:"column:display_name_char;type:varchar;"` // compute field display_name_char store is False
	AliasName       string `json:"alias_name" gorm:"column:alias_name;type:varchar;"`
	Abbreviation    string `json:"abbreviation" gorm:"column:abbreviation;type:varchar;"`
	BlankFields     string `json:"blank_fields" gorm:"column:blank_fields;type:varchar;"` // compute field blank_fields store is False
	//Unit   *ResCurrency `json:"unit" gorm:"column:unit;type:int;"`  // compute field unit store is False
	ParentRef        string `json:"parent_ref" gorm:"column:parent_ref;type:varchar;"`                 // compute field parent_ref store is False
	DepartmentIdName string `json:"department_id_name" gorm:"column:department_id_name;type:varchar;"` // compute field department_id_name store is False
	//Gender   None `json:"gender" gorm:"column:gender;type:;"`  // compute field gender store is False
	//RegionsId   *ResCountryStateRegions `json:"regions_id" gorm:"column:regions_id;type:int;"`  // compute field regions_id store is True
	Wechat      string `json:"wechat" gorm:"column:wechat;type:varchar;"`
	Qq          string `json:"qq" gorm:"column:qq;type:varchar;"`
	Skype       string `json:"skype" gorm:"column:skype;type:varchar;"`
	NormalState int    `json:"normal_state" gorm:"column:normal_state;type:int;not null"`
	State       int    `json:"state" gorm:"column:state;type:int;not null"` // compute field state store is True
	//PhoneCountryCode   *CountryCodeList `json:"phone_country_code" gorm:"column:phone_country_code;type:int;"`  // compute field phone_country_code store is False
	PhoneAreaCode string `json:"phone_area_code" gorm:"column:phone_area_code;type:varchar;"` // compute field phone_area_code store is False
	PhoneText     string `json:"phone_text" gorm:"column:phone_text;type:varchar;"`           // compute field phone_text store is False
	//MobileCountryCode   *CountryCodeList `json:"mobile_country_code" gorm:"column:mobile_country_code;type:int;"`  // compute field mobile_country_code store is False
	MobileText string `json:"mobile_text" gorm:"column:mobile_text;type:varchar;"` // compute field mobile_text store is False
	//DefaultCurrencyId   *ResCurrency `json:"default_currency_id" gorm:"column:default_currency_id;type:int;"`
	//DepartmentId   *HrDepartment `json:"department_id" gorm:"column:department_id;type:int;"`
	EndDate                time.Time   `json:"end_date" gorm:"column:end_date;type:timestamp(0) with time zone;"`
	Scale                  int         `json:"scale" gorm:"column:scale;type:int;"`
	OutUserId              int         `json:"out_user_id" gorm:"column:out_user_id;type:int;"`
	OutSbuId               int         `json:"out_sbu_id" gorm:"column:out_sbu_id;type:int;"`
	OutContactId           int         `json:"out_contact_id" gorm:"column:out_contact_id;type:int;"`
	OutAddressId           int         `json:"out_address_id" gorm:"column:out_address_id;type:int;"`
	ContactType            int         `json:"contact_type" gorm:"column:contact_type;type:int;"`
	IsUser                 bool        `json:"is_user" gorm:"column:is_user;type:bool;"`
	IsDepartment           bool        `json:"is_department" gorm:"column:is_department;type:bool;"`
	ContactState           int         `json:"contact_state" gorm:"column:contact_state;type:int;"`
	Level                  int         `json:"level" gorm:"column:level;type:int;"`
	IManageAndParent       bool        `json:"i_manage_and_parent" gorm:"column:i_manage_and_parent;type:bool;"` // compute field i_manage_and_parent store is False
	IsMainAccount          bool        `json:"is_main_account" gorm:"column:is_main_account;type:bool;"`         // compute field is_main_account store is False
	SbuId                  *ResPartner `json:"sbu_id" gorm:"column:sbu_id;type:int;"`
	InvitationId           *ResPartner `json:"invitation_id" gorm:"column:invitation_id;type:int;"`
	Ishomepage             bool        `json:"ishomepage" gorm:"column:ishomepage;type:bool;"`
	PurchaseOrderLineCount int         `json:"purchase_order_line_count" gorm:"column:purchase_order_line_count;type:int;"` // compute field purchase_order_line_count store is False
	SaleOrderLineCount     int         `json:"sale_order_line_count" gorm:"column:sale_order_line_count;type:int;"`         // compute field sale_order_line_count store is False
	//ImportLines   *[]FileImport `json:"import_lines" gorm:"column:import_lines;type:;"`
	DateOfDeliveryHk int     `json:"date_of_delivery_hk" gorm:"column:date_of_delivery_hk;type:int;"`
	DateOfDeliverySz int     `json:"date_of_delivery_sz" gorm:"column:date_of_delivery_sz;type:int;"`
	MinOrderQuantity float32 `json:"min_order_quantity" gorm:"column:min_order_quantity;type:;"`
	SalePercentPrice float32 `json:"sale_percent_price" gorm:"column:sale_percent_price;type:;"`
	//RebatePricelistId   *ProductPricelist `json:"rebate_pricelist_id" gorm:"column:rebate_pricelist_id;type:int;"`
	//PurchasePricelistId   *ProductPricelist `json:"purchase_pricelist_id" gorm:"column:purchase_pricelist_id;type:int;"`
}

func (parent ResPartner) Create(value interface{}) {
	db.PG.Create(value)
}
