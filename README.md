Inside the main directory, create a folder for the module. The name of the folder should match the name of the package, at the top of the .go files within.

normal test files use the name of the file with _test. They do not import the package. It should not be necessary to namespace the functions, etc., with the package name.

example_test.go files can be helpful to test output. We do import the pkg. We do have to namespace the functions, etc.