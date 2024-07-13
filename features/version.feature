Feature: Version
    In order to be informed about the application
    As a user
    I need to be able to access the version information

    Scenario: Validating the version
        When I execute the version command
        Then I should see the version from the VERSION file
        And the exit code should be 0
