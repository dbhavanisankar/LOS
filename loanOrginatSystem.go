/*
CreateProspect
GetDetailsByProspectId
ViewProspect
CreateApplication
GetDetailsByApplicantId
ViewApplication
ViewProperties
GetDetailsByPropertyId
UpdateValuerByPropertyId
UpdateSolicitorByPropertyId
UpdateUnderwriterByApplicantId
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// LoanOriginatingSystem will implement the processes
type LoanOriginatingSystem struct {
}

type customerInfo struct {
	CustomerID						string `json:"customerID"`
	ProspectNO                		string `json:"prospectNO"`

	FirstName						string `json:"firstName"`
	LastName     					string `json:"lastName"`
	DOB				 				string `json:"dob"`
	Address     					string `json:"address"`
	SSN				    			string `json:"ssn"`
	PassportNo          			string `json:"passportNo"`
	PurposeOfLoan       			string `json:"purposeOfLoan"`
	ContactNumber       			string `json:"contactNumber"`
	EmailID           				string `json:"emailID"`
	LoanType						string `json:"loanType"`

	AssetCash     					string `json:"assetCash"`
	AssetChecking_SavingsAccount   	string `json:"assetChecking_SavingsAccount"`
	AssetStocks     				string `json:"assetStocks"`
	AssetRetirementAccount			string `json:"assetRetirementAccount"`
	AssetSurrenderValueOfPolicies	string `json:"assetSurrenderValueOfPolicies"`
	AssetSum		            	string `json:"assetSum"`

	LiabilitiesMortgage		        string `json:"liabilitiesMortgage"`
	LiabilitiesCreditcards			string `json:"liabilitiesCreditcards"`
	LiabilitiesOtherLoanAmounts	    string `json:"liabilitiesOtherLoanAmounts"`
	LiabilitiesAlimony_ChildSupport string `json:"liabilitiesAlimony_ChildSupport"`
	LiabilitiesNetWorth     		string `json:"liabilitiesNetworth"`
}
type LoanTransactions struct {
	ProspectNO                		string `json:"prospectNO"`
	ApplicantID              		string `json:"applicantID"`
	PropertyID           			string `json:"propertyID"`
	ApplicationStatus				string `json:"applicationStatus"`

	Employment_CurrentOccupation    string `json:"employmentCurrentOccupation"`
	Employment_SalariedEmployerName string `json:"employment_SalariedEmployerName"`
	Employment_BusinessType 		string `json:"employment_BusinessType"`
	Employment_MothlyIncome 		string `json:"employment_MothlyIncome"`

	Loan_BorrowerAddress		    string `json:"loan_BorrowerAddress"`
	Loan_Purpose					string `json:"loan_Purpose"`
	Loan_Amount					    string `json:"loan_Amount"`
	Loan_Term   					string `json:"loan_Term"`

	Property_Description			string `json:"property_Description"`
	Property_Type 					string `json:"property_Type"`
	Property_PurchaseDate      		string `json:"property_PurchaseDate"`
	Property_BuiltYear           	string `json:"property_BuiltYear"`
	Property_Tenure					string `json:"property_Tenure"`
	Property_ExLocalAuthority		string `json:"property_ExLocalAuthority"`
	Property_NoOfBedrooms			string `json:"property_NoOfBedrooms"`
	Property_NoOfBathrooms 			string `json:"property_NoOfBathrooms"`
	Property_NoOfReceptionRooms		string `json:"property_NoOfReceptionRooms"`
	Property_RentalIncome           string `json:"property_RentalIncome"`
	Property_CurrentlyLetStatus		string `json:"property_CurrentlyLetStatus"`

	DocumentHash    				string `json:"documentHash"`
	ApplicationSentToValuerStatus   string `json:"applicationSentToValuerStatus"`
	ApplicationSentToLawyerStatus   string `json:"applicationSentToLawyerStatus"`

	/*Property related*/
	Valuer_YearOfConstruction       string `json:"valuer_YearOfConstruction"`
	Valuer_PropertyTitle 			string `json:"valuer_PropertyTitle"`
	Valuer_BuilderName     			string `json:"valuer_BuilderName"`
	Valuer_RERARegnNo   			string `json:"valuer_RERARegnNo"`
	Valuer_BuilderPANCard     		string `json:"valuer_BuilderPANCard"`
	Valuer_YearofPurchase     		string `json:"valuer_YearofPurchase"`
	Valuer_PropertyDocuments        string `json:"valuer_PropertyDocuments"`
	Valuer_LandParcelArea      		string `json:"valuer_LandParcelArea"`
	Valuer_MunicipalLimitations 	string `json:"valuer_MunicipalLimitations"`
	Valuer_LotAttributes    		string `json:"valuer_LotAttributes"`
	Valuer_ConstructionArea   		string `json:"valuer_ConstructionArea"`
	Valuer_StreetFrontage     		string `json:"valuer_StreetFrontage"`
	Valuer_Presentation_Layout     	string `json:"valuer_Presentation_Layout"`
	Valuer_AccessObstructions       string `json:"valuer_AccessObstructions"`
	Valuer_OffStreetParking      	string `json:"valuer_OffStreetParking"`
	Valuer_FutureObstructions 		string `json:"valuer_FutureObstructions"`
	Valuer_ValuationAmount    		string `json:"valuer_ValuationAmount"`
	Valuer_FSIGranted   			string `json:"valuer_FSIGranted"`
	Valuer_OtherFactors     		string `json:"valuer_OtherFactors"`
	/*Apartment related*/
	Valuer_ApartmentSize            string `json:"valuer_ApartmentSize"`
	Valuer_Bedrooms     			string `json:"valuer_Bedrooms"`
	Valuer_Bathrooms   				string `json:"valuer_Bathrooms"`
	Valuer_ConstructionCondition    string `json:"valuer_ConstructionCondition"`
	Valuer_KitchenSize			    string `json:"valuer_KitchenSize"`
	Valuer_ApartmentValuation		string `json:"valuer_ApartmentValuation"`

	Lawyer_PropertyTitle        	string `json:"lawyer_PropertyTitle"`
	Lawyer_PropertyTitleComment     string `json:"lawyer_PropertyTitleComment"`

	Lawyer_PermissionFromMunicipal 	string `json:"lawyer_PermissionFromMunicipal"`
	Lawyer_PermissionFromMunicipalComment	string `json:"lawyer_PermissionFromMunicipalComment"`

	Lawyer_ClearPastTitle    		string `json:"lawyer_ClearPastTitle"`
	Lawyer_ClearPastTitleComment  	string `json:"lawyer_ClearPastTitleComment"`

	Lawyer_LitigationsPending   	string `json:"lawyer_LitigationsPending"`
	Lawyer_LitigationsPendingComment	  string `json:"lawyer_LitigationsPendingComment"`

	Lawyer_LitigationByBuilder     	string `json:"lawyer_LitigationByBuilder"`
	Lawyer_LitigationByBuildercomment	  string `json:"lawyer_LitigationByBuilderComment"`

	Lawyer_ProvenanceReference1     string `json:"lawyer_ProvenanceReference1"`
	Lawyer_ProvenanceReference1Comment    string `json:"lawyer_ProvenanceReference1Comment"`

	Lawyer_ProvenanceReference2     string `json:"lawyer_ProvenanceReference2"`
	Lawyer_ProvenanceReference2Comment    string `json:"lawyer_ProvenanceReference2Comment"`

	Lawyer_ProvenanceReference3     string `json:"lawyer_ProvenanceReference3"`
	Lawyer_ProvenanceReference3Comment    string `json:"lawyer_ProvenanceReference3Comment"`

	Lawyer_ClientAgreement_Apartment	  string `json:"lawyer_ClientAgreement_Apartment"`
	Lawyer_ApprovalStatus 			string `json:"lawyer_ApprovalStatus"`
	Lawyer_ApprovalReason    		string `json:"lawyer_ApprovalReason"`

	Underwriter_ProductVerified    		string `json:"underwriter_ProductVerified"`
	Underwriter_PropertyVerified   		string `json:"underwriter_PropertyVerified"`
	Underwriter_EmployeeDetailsVerified string `json:"underwriter_EmployeeDetailsVerified"`
	Underwriter_MothlyExpenseVerified   string `json:"underwriter_MothlyExpenseVerified"`
	Underwriter_AssetLiabilityVerified	string `json:"underwriter_AssetLiabilityVerified"`
	Underwriter_ApprovalStatus 		    string `json:"underwriter_ApprovalStatus"`
	Underwriter_ApprovalReason			string `json:"underwriter_ApprovalReason"`
}


//invoke methods CreateProspect
func (t *LoanOriginatingSystem) CreateProspect(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var objcustomerInfo customerInfo
	var objLoanTransactions LoanTransactions
	var err error
	
	fmt.Println("Entering CreateProspect")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	fmt.Println("Args [0] is : %v\n", args[0])

	//unmarshal customerInfo data from UI to "customerInfo" struct
	err = json.Unmarshal([]byte(args[0]), &objcustomerInfo)
	if err != nil {
		fmt.Printf("Unable to unmarshal CreateProspect input customerInfo: %s\n", err)
		return shim.Error(err.Error())
		}

	fmt.Println("customerInfo object Customer ID variable value is : %s\n", objcustomerInfo.CustomerID)

	// Data insertion for Couch DB starts here CustomerID
	transJSONasBytes, err := json.Marshal(objcustomerInfo)
	err = stub.PutState(objcustomerInfo.CustomerID, transJSONasBytes)
	// Data insertion for Couch DB ends here

	//unmarshal LoanTransactions data from UI to "LoanTransactions" struct
	err = json.Unmarshal([]byte(args[0]), &objLoanTransactions)
	if err != nil {
		fmt.Printf("Unable to unmarshal CreateProspect input customerInfo: %s\n", err)
		return shim.Error(err.Error())
		}

	fmt.Println("customerInfo object Customer ID variable value is : %s\n", objLoanTransactions.ProspectNO)

	// Data insertion for Couch DB starts here 
	transJSONasBytesLoan, err := json.Marshal(objLoanTransactions)
	err = stub.PutState(objLoanTransactions.ProspectNO, transJSONasBytesLoan)
	// Data insertion for Couch DB ends here

	fmt.Println("Create CustomerInfo Successfully Done")

	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
		//return nil,nil
	}
	return shim.Success(nil)
}

//GetDetailsByProspectId
func (t *LoanOriginatingSystem) GetDetailsByProspectId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	fmt.Println("Entering GetDetailsByProspectId")

	var prospectNo = args[0]
	//fetch data from couch db starts here
	queryString := fmt.Sprintf("{\"selector\":{\"prospectNO\":{\"$eq\": \"%s\"}}}", prospectNo)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here

	if err != nil {
		fmt.Printf("Unable to read the customer details : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Printf("list of customer details based on prospectId: %v\n", queryResults)
	return shim.Success(queryResults)
}

//ViewProspect
func (t *LoanOriginatingSystem) ViewProspect(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	fmt.Println("Entering ViewProspect")

	//fetch data from couch db starts here
	queryString := fmt.Sprintf("{\"selector\":{\"prospectNO\":{\"$ne\": \"null\"}}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here
	if err != nil {
		fmt.Printf("Unable to read the list of prospects: %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Printf("list of all prospects details : %v\n", queryResults)
	return shim.Success(queryResults)

}
/*
//invoke methods CreateApplication
func (t *LoanOriginatingSystem) CreateApplication(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var objLoanTransactions LoanTransactions
	var err error
	
	fmt.Println("Entering CreateApplication")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	fmt.Println("Args [0] is : %v\n", args[0])

	//unmarshal LoanTransactions data from UI to "LoanTransactions" struct
	err = json.Unmarshal([]byte(args[0]), &objLoanTransactions)
	if err != nil {
		fmt.Printf("Unable to unmarshal CreateApplication input customerInfo: %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("customerInfo object ProspectNO variable value is : %s\n", objLoanTransactions.ProspectNO)

	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(objLoanTransactions)
	err = stub.PutState(objLoanTransactions.ProspectNO, transJSONasBytes)
	// Data insertion for Couch DB ends here

	fmt.Println("Create CreateApplication Successfully Done")

	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
*/
//UpdateSolicitorByPropertyId
func (t *LoanOriginatingSystem) CreateApplication(stub shim.ChaincodeStubInterface, args1 []string) pb.Response {
	var objUILoanTransactions LoanTransactions
	var objBCLoanTransactions LoanTransactions
	var err error

	fmt.Println("Entering CreateApplication")

	if len(args1) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args1[0]), &objUILoanTransactions)
	if err != nil {
		fmt.Printf("Unable to marshal  CreateApplication input CreateApplication : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n refno PropertyID is : ", objUILoanTransactions.ProspectNO)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(objUILoanTransactions.ProspectNO)
	err = json.Unmarshal(bytesread, &objBCLoanTransactions)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\nobjBCLoanTransactions in updateLawyer : %s ", objBCLoanTransactions)
	
	objBCLoanTransactions.ProspectNO = objUILoanTransactions.ProspectNO
	objBCLoanTransactions.ApplicantID = objUILoanTransactions.ApplicantID
	objBCLoanTransactions.PropertyID = objUILoanTransactions.PropertyID
	objBCLoanTransactions.ApplicationStatus = objUILoanTransactions.ApplicationStatus

	objBCLoanTransactions.Employment_CurrentOccupation = objUILoanTransactions.Employment_CurrentOccupation
	objBCLoanTransactions.Employment_SalariedEmployerName = objUILoanTransactions.Employment_SalariedEmployerName
	objBCLoanTransactions.Employment_BusinessType = objUILoanTransactions.Employment_BusinessType
	objBCLoanTransactions.Employment_MothlyIncome = objUILoanTransactions.Employment_MothlyIncome

	objBCLoanTransactions.Loan_BorrowerAddress = objUILoanTransactions.Loan_BorrowerAddress
	objBCLoanTransactions.Loan_Purpose = objUILoanTransactions.Loan_Purpose
	objBCLoanTransactions.Loan_Amount = objUILoanTransactions.Loan_Amount
	objBCLoanTransactions.Loan_Term = objUILoanTransactions.Loan_Term

	objBCLoanTransactions.Property_Description = objUILoanTransactions.Property_Description
	objBCLoanTransactions.Property_Type = objUILoanTransactions.Property_Type
	objBCLoanTransactions.Property_PurchaseDate = objUILoanTransactions.Property_PurchaseDate
	objBCLoanTransactions.Property_BuiltYear = objUILoanTransactions.Property_BuiltYear
	objBCLoanTransactions.Property_Tenure = objUILoanTransactions.Property_Tenure
	objBCLoanTransactions.Property_ExLocalAuthority = objUILoanTransactions.Property_ExLocalAuthority
	objBCLoanTransactions.Property_NoOfBedrooms = objUILoanTransactions.Property_NoOfBedrooms
	objBCLoanTransactions.Property_NoOfBathrooms = objUILoanTransactions.Property_NoOfBathrooms
	objBCLoanTransactions.Property_NoOfReceptionRooms = objUILoanTransactions.Property_NoOfReceptionRooms
	objBCLoanTransactions.Property_RentalIncome = objUILoanTransactions.Property_RentalIncome
	objBCLoanTransactions.Property_CurrentlyLetStatus = objUILoanTransactions.Property_CurrentlyLetStatus

	objBCLoanTransactions.DocumentHash = objUILoanTransactions.DocumentHash
	objBCLoanTransactions.ApplicationSentToValuerStatus = objUILoanTransactions.ApplicationSentToValuerStatus
	objBCLoanTransactions.ApplicationSentToLawyerStatus = objUILoanTransactions.ApplicationSentToLawyerStatus

	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(objBCLoanTransactions)
	err = stub.PutState(objUILoanTransactions.ProspectNO, transJSONasBytes)
	// Data insertion for Couch DB ends here

	fmt.Println("Create Application transaction Successfully updated. Details updated in LoanTransactions struct")

	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//GetDetailsByApplicantId
func (t *LoanOriginatingSystem) GetDetailsByApplicantId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	fmt.Println("Entering GetDetailsByApplicantId")

	var applicantId = args[0]
	//fetch data from couch db starts here
	queryString := fmt.Sprintf("{\"selector\":{\"applicantID\":{\"$eq\": \"%s\"}}}", applicantId)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here

	if err != nil {
		fmt.Printf("Unable to read the applicant details : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Printf("list of applicant details based on applicant id: %v\n", queryResults)
	return shim.Success(queryResults)
}

//ViewApplication
func (t *LoanOriginatingSystem) ViewApplication(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	fmt.Println("Entering ViewApplication")

	//fetch data from couch db starts here
	//queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"customerID\":{\"$ne\":\"null\"}},{\"applicantID\":{\"$ne\":\"null\"}}]}}")
	queryString := fmt.Sprintf("{\"selector\":{\"applicantID\":{\"$ne\": \"null\"}}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here
	if err != nil {
		fmt.Printf("Unable to read the list of prospects: %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Printf("list of all prospects details : %v\n", queryResults)
	return shim.Success(queryResults)
}

//ViewProperties
func (t *LoanOriginatingSystem) ViewProperties(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	fmt.Println("Entering ViewProperties")

	//fetch data from couch db starts here
	queryString := fmt.Sprintf("{\"selector\":{\"propertyID\":{\"$ne\": \"null\"}}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here

	if err != nil {
		fmt.Printf("Unable to read the list of Properties: %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Printf("list of all Properties details : %v\n", queryResults)
	return shim.Success(queryResults)
}

//GetDetailsByPropertyId
func (t *LoanOriginatingSystem) GetDetailsByPropertyId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	fmt.Println("Entering GetDetailsByPropertyId")

	var propertyId = args[0]
	//fetch data from couch db starts here
	queryString := fmt.Sprintf("{\"selector\":{\"propertyID\":{\"$eq\": \"%s\"}}}", propertyId)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here

	if err != nil {
		fmt.Printf("Unable to read the property details : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Printf("list of property details based on property id: %v\n", queryResults)
	return shim.Success(queryResults)
}

//UpdateValuerByPropertyId
func (t *LoanOriginatingSystem) UpdateValuerByPropertyId(stub shim.ChaincodeStubInterface, args1 []string) pb.Response {
	var objUILoanTransactions LoanTransactions
	var objBCLoanTransactions LoanTransactions
	var err error

	fmt.Println("Entering UpdateValuerByPropertyId")

	if len(args1) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args1[0]), &objUILoanTransactions)
	if err != nil {
		fmt.Printf("Unable to marshal  CreateApplication input UpdateValuerByPropertyId : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n refno PropertyID is : ", objUILoanTransactions.ProspectNO)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(objUILoanTransactions.ProspectNO)
	err = json.Unmarshal(bytesread, &objBCLoanTransactions)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\nobjBCLoanTransactions in updateValuer : %s ", objBCLoanTransactions)
	//prooperty related
	objBCLoanTransactions.ProspectNO = objUILoanTransactions.ProspectNO
	objBCLoanTransactions.ApplicationStatus = objUILoanTransactions.ApplicationStatus
	objBCLoanTransactions.Valuer_YearOfConstruction = objUILoanTransactions.Valuer_YearOfConstruction
	objBCLoanTransactions.Valuer_PropertyTitle = objUILoanTransactions.Valuer_PropertyTitle
	objBCLoanTransactions.Valuer_BuilderName = objUILoanTransactions.Valuer_BuilderName
	objBCLoanTransactions.Valuer_RERARegnNo = objUILoanTransactions.Valuer_RERARegnNo
	objBCLoanTransactions.Valuer_BuilderPANCard = objUILoanTransactions.Valuer_BuilderPANCard
	objBCLoanTransactions.Valuer_YearofPurchase = objUILoanTransactions.Valuer_YearofPurchase
	objBCLoanTransactions.Valuer_PropertyDocuments = objUILoanTransactions.Valuer_PropertyDocuments
	objBCLoanTransactions.Valuer_LandParcelArea = objUILoanTransactions.Valuer_LandParcelArea
	objBCLoanTransactions.Valuer_MunicipalLimitations = objUILoanTransactions.Valuer_MunicipalLimitations
	objBCLoanTransactions.Valuer_LotAttributes = objUILoanTransactions.Valuer_LotAttributes
	objBCLoanTransactions.Valuer_ConstructionArea = objUILoanTransactions.Valuer_ConstructionArea
	objBCLoanTransactions.Valuer_StreetFrontage = objUILoanTransactions.Valuer_StreetFrontage
	objBCLoanTransactions.Valuer_Presentation_Layout = objUILoanTransactions.Valuer_Presentation_Layout
	objBCLoanTransactions.Valuer_AccessObstructions = objUILoanTransactions.Valuer_AccessObstructions
	objBCLoanTransactions.Valuer_OffStreetParking = objUILoanTransactions.Valuer_OffStreetParking
	objBCLoanTransactions.Valuer_FutureObstructions = objUILoanTransactions.Valuer_FutureObstructions
	objBCLoanTransactions.Valuer_ValuationAmount = objUILoanTransactions.Valuer_ValuationAmount
	objBCLoanTransactions.Valuer_FSIGranted = objUILoanTransactions.Valuer_FSIGranted
	objBCLoanTransactions.Valuer_OtherFactors = objUILoanTransactions.Valuer_OtherFactors
	//apartment related
	objBCLoanTransactions.Valuer_ApartmentSize = objUILoanTransactions.Valuer_ApartmentSize
	objBCLoanTransactions.Valuer_Bedrooms = objUILoanTransactions.Valuer_Bedrooms
	objBCLoanTransactions.Valuer_Bathrooms = objUILoanTransactions.Valuer_Bathrooms
	objBCLoanTransactions.Valuer_ConstructionCondition = objUILoanTransactions.Valuer_ConstructionCondition
	objBCLoanTransactions.Valuer_KitchenSize = objUILoanTransactions.Valuer_KitchenSize
	objBCLoanTransactions.Valuer_ApartmentValuation = objUILoanTransactions.Valuer_ApartmentValuation
	

	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(objBCLoanTransactions)
	err = stub.PutState(objUILoanTransactions.ProspectNO, transJSONasBytes)
	// Data insertion for Couch DB ends here

	fmt.Println("Valuer transaction Successfully updated. Details updated in LoanTransactions struct")

	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//UpdateSolicitorByPropertyId
func (t *LoanOriginatingSystem) UpdateSolicitorByPropertyId(stub shim.ChaincodeStubInterface, args1 []string) pb.Response {
	var objUILoanTransactions LoanTransactions
	var objBCLoanTransactions LoanTransactions
	var err error

	fmt.Println("Entering UpdateSolicitorByPropertyId")

	if len(args1) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args1[0]), &objUILoanTransactions)
	if err != nil {
		fmt.Printf("Unable to marshal  CreateApplication input UpdateSolicitorByPropertyId : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n refno PropertyID is : ", objUILoanTransactions.ProspectNO)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(objUILoanTransactions.ProspectNO)
	err = json.Unmarshal(bytesread, &objBCLoanTransactions)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\nobjBCLoanTransactions in updateLawyer : %s ", objBCLoanTransactions)
	
	objBCLoanTransactions.ProspectNO = objUILoanTransactions.ProspectNO
	objBCLoanTransactions.ApplicationStatus = objUILoanTransactions.ApplicationStatus
	objBCLoanTransactions.Lawyer_PropertyTitle = objUILoanTransactions.Lawyer_PropertyTitle
	objBCLoanTransactions.Lawyer_PropertyTitleComment = objUILoanTransactions.Lawyer_PropertyTitleComment
	objBCLoanTransactions.Lawyer_PermissionFromMunicipal = objUILoanTransactions.Lawyer_PermissionFromMunicipal
	objBCLoanTransactions.Lawyer_PermissionFromMunicipalComment = objUILoanTransactions.Lawyer_PermissionFromMunicipalComment
	objBCLoanTransactions.Lawyer_ClearPastTitle = objUILoanTransactions.Lawyer_ClearPastTitle
	objBCLoanTransactions.Lawyer_ClearPastTitleComment = objUILoanTransactions.Lawyer_ClearPastTitleComment
	objBCLoanTransactions.Lawyer_LitigationsPending = objUILoanTransactions.Lawyer_LitigationsPending
	objBCLoanTransactions.Lawyer_LitigationsPendingComment = objUILoanTransactions.Lawyer_LitigationsPendingComment
	objBCLoanTransactions.Lawyer_LitigationByBuilder = objUILoanTransactions.Lawyer_LitigationByBuilder
	objBCLoanTransactions.Lawyer_LitigationByBuildercomment = objUILoanTransactions.Lawyer_LitigationByBuildercomment
	objBCLoanTransactions.Lawyer_ProvenanceReference1 = objUILoanTransactions.Lawyer_ProvenanceReference1
	objBCLoanTransactions.Lawyer_ProvenanceReference1Comment = objUILoanTransactions.Lawyer_ProvenanceReference1Comment
	objBCLoanTransactions.Lawyer_ProvenanceReference2 = objUILoanTransactions.Lawyer_ProvenanceReference2
	objBCLoanTransactions.Lawyer_ProvenanceReference2Comment = objUILoanTransactions.Lawyer_ProvenanceReference2Comment
	objBCLoanTransactions.Lawyer_ProvenanceReference3 = objUILoanTransactions.Lawyer_ProvenanceReference3
	objBCLoanTransactions.Lawyer_ProvenanceReference3Comment = objUILoanTransactions.Lawyer_ProvenanceReference3Comment
	objBCLoanTransactions.Lawyer_ClientAgreement_Apartment = objUILoanTransactions.Lawyer_ClientAgreement_Apartment
	objBCLoanTransactions.Lawyer_ApprovalStatus = objUILoanTransactions.Lawyer_ApprovalStatus
	objBCLoanTransactions.Lawyer_ApprovalReason = objUILoanTransactions.Lawyer_ApprovalReason
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(objBCLoanTransactions)
	err = stub.PutState(objUILoanTransactions.ProspectNO, transJSONasBytes)
	// Data insertion for Couch DB ends here

	fmt.Println("Lawyer transaction Successfully updated. Details updated in LoanTransactions struct")

	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
//UpdateUnderwriterByApplicantId
func (t *LoanOriginatingSystem) UpdateUnderwriterByApplicantId(stub shim.ChaincodeStubInterface, args1 []string) pb.Response {
	var objUILoanTransactions LoanTransactions
	var objBCLoanTransactions LoanTransactions
	var err error

	fmt.Println("Entering UpdateUnderwriterByApplicantId")

	if len(args1) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args1[0]), &objUILoanTransactions)
	if err != nil {
		fmt.Printf("Unable to marshal  CreateApplication input UpdateUnderwriterByApplicantId : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n refno ApplicantID is : ", objUILoanTransactions.ProspectNO)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(objUILoanTransactions.ProspectNO)
	err = json.Unmarshal(bytesread, &objBCLoanTransactions)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\nobjBCLoanTransactions in update Underwriter : %s ", objBCLoanTransactions)
	
	objBCLoanTransactions.ProspectNO = objUILoanTransactions.ProspectNO
	objBCLoanTransactions.ApplicationStatus = objUILoanTransactions.ApplicationStatus
	objBCLoanTransactions.Underwriter_ProductVerified = objUILoanTransactions.Underwriter_ProductVerified
	objBCLoanTransactions.Underwriter_PropertyVerified = objUILoanTransactions.Underwriter_PropertyVerified
	objBCLoanTransactions.Underwriter_EmployeeDetailsVerified = objUILoanTransactions.Underwriter_EmployeeDetailsVerified
	objBCLoanTransactions.Underwriter_MothlyExpenseVerified = objUILoanTransactions.Underwriter_MothlyExpenseVerified
	objBCLoanTransactions.Underwriter_AssetLiabilityVerified = objUILoanTransactions.Underwriter_AssetLiabilityVerified
	objBCLoanTransactions.Underwriter_ApprovalStatus = objUILoanTransactions.Underwriter_ApprovalStatus
	objBCLoanTransactions.Underwriter_ApprovalReason = objUILoanTransactions.Underwriter_ApprovalReason
		
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(objBCLoanTransactions)
	err = stub.PutState(objUILoanTransactions.ProspectNO, transJSONasBytes)
	// Data insertion for Couch DB ends here

	fmt.Println("Underwriter transaction Successfully updated. Details updated in LoanTransactions struct")

	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}


// getQueryResultForQueryString
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// Init sets up the chaincode
func (t *LoanOriginatingSystem) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Initiate the chaincde")
	return shim.Success(nil)
	//	return nil,nil
}

// Invoke the function in the chaincode
func (t *LoanOriginatingSystem) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "CreateProspect" {
		return t.CreateProspect(stub, args)
	}
	if function == "CreateApplication" {
		return t.CreateApplication(stub, args)
	}
	if function == "UpdateValuerByPropertyId" {
		return t.UpdateValuerByPropertyId(stub, args)
	}
	if function == "UpdateSolicitorByPropertyId" {
		return t.UpdateSolicitorByPropertyId(stub, args)
	}
	if function == "UpdateUnderwriterByApplicantId" {
		return t.UpdateUnderwriterByApplicantId(stub, args)
	}
	if function == "ViewProspect" {
		return t.ViewProspect(stub, args)
	}
	if function == "GetDetailsByProspectId" {
		return t.GetDetailsByProspectId(stub, args)
	}
	if function == "GetDetailsByApplicantId" {
		return t.GetDetailsByApplicantId(stub, args)
	}
	if function == "ViewApplication" {
		return t.ViewApplication(stub, args)
	}
	if function == "ViewProperties" {
		return t.ViewProperties(stub, args)
	}
	if function == "GetDetailsByPropertyId" {
		return t.GetDetailsByPropertyId(stub, args)
	}
	
	fmt.Println("Function not found")
	return shim.Error("Received unknown function invocation")
	//return nil, nil
}

func main() {
	err := shim.Start(new(LoanOriginatingSystem))
	if err != nil {
		fmt.Println("Could not start Chaincode")
	} else {
		fmt.Println("Chaincode successfully started")
	}

}
