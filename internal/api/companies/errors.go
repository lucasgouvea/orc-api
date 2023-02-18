package companies

import "errors"

var InvalidCompanyTypeErr = errors.New("Invalid company type.")
var MissingIntermediatedErr = errors.New("AGGREGATE company must have intermediateds.")
var ContractIntermediatedErr = errors.New("CONTRACT company must not have intermediateds.")
