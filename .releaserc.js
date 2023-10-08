module.exports = {
    branches: ['main'], // Define your default branch name

    plugins: [
        '@semantic-release/commit-analyzer', // Analyzes commit messages to determine the next version.
        '@semantic-release/release-notes-generator', // Generates release notes from commit messages.
        '@semantic-release/changelog', // Updates a CHANGELOG.md file based on the release notes.
        '@semantic-release/github', // Publishes releases on GitHub.
        '@semantic-release/git' // Commits version bump and changelog files.
    ],

    prepare: [
        {
            path: '@semantic-release/exec',
            cmd: 'make build' // Build your Go application here if needed.
        }
    ]

    // Customize the release process further if necessary.
}
