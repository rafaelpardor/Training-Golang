package stringutil

// If the variable start with a lowercase, this will not be exported.
var myName = "Variable that start with a lowercase but is called by another variable that start with uppercase."

// Intead, it start with a Uppercase, this will be exported.
var Exported = myName
