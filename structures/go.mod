module project

go 1.17

require internal/customer v1.0.0
replace internal/customer => ./internal/customer
require internal/debtor v1.0.0
replace internal/debtor => ./internal/debtor
