Feature: New Command
  In order to create new resources
  As a user
  I need to be able to create new resources

  Scenario: Create a new project
    Given a module path "github.com/ocrosby/project1"
    And a target directory "~/go/github.com/ocrosby/project1"
    When I create a new project
    Then there should be no errors
    And I should see a