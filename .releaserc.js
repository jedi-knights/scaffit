module.exports = {
    branches: ['main'],
    ci: true,
    debug: true,
    dryRun: false,
    tagFormat: '${version}',

    // Global plugin options (will be passed to all plugins)
    preset: 'conventionalcommits',

    // Reponsible for verifying conditions necessary to proceed with the release
    verifyConditions: [
        '@semantic-release/changelog',
        '@semantic-release/git',
        '@semantic-release/github'
    ],

    // Responsible for determining the type of release to make (major, minor or patch)
    // See: https://github.com/semantic-release/commit-analyzer#configuration
    analyzeCommits: [
        '@semantic-release/commit-analyzer'
    ],

    // Responsible for generating the content of the release note
    generateNotes: [
        '@semantic-release/release-notes-generator'
    ],

    // Responsible for preparing the release, for example creating or updating files such as package.json, CHANGELOG.md, documentation or compiled assets
    prepare: [
        '@semantic-release/changelog',
        '@semantic-release/git'
    ],

    // Responsible for preparing the release, for example creating or updating files such as package.json, CHANGELOG.md, documentation or compiled assets
    publish: [
        '@semantic-release/github'
    ],

    plugins: [
        [
            '@semantic-release/commit-analyzer',
            {
                preset: 'conventionalcommits',
                releaseRules: [
                    { type: 'docs', release: 'patch' },
                    { type: 'refactor', release: 'patch' },
                    { type: 'style', release: 'patch' }
                ],
                parserOpts: {
                    noteKeywords: ['BREAKING CHANGE', 'BREAKING CHANGES', 'BREAKING']
                }
            }
        ],

        // for each release, the CHANGELOG.md file will be created or updated.
        '@semantic-release/release-notes-generator',
        [
            '@semantic-release/changelog',
            {
                changelogFile: 'CHANGELOG.md'
            }
        ],
        [
            '@semantic-release/git',
            {
                assets: ['CHANGELOG.md']
            }
        ],

        '@semantic-release/changelog',
        [
            '@semantic-release/exec',
            {
                prepareCmd: 'echo ${nextRelease.version} > VERSION'
            }
        ],

        [
            // semantic-release plugin to commit release assets to the project's git repository.
            // https://github.com/semantic-release/git
            //
            // The Git user associated with the Git credentials has to be able to push commit to the release branch.
            '@semantic-release/git',
            {
                assets: ['CHANGELOG.md', 'VERSION'],
                //message: 'chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}'
                message: 'chore(release): ${nextRelease.version} [skip ci]'
            }
        ],

        [
            // semantic-release plugin to publish a GitHub release and comment on released pull requests/issues.
            // https://github.com/semantic-release/github
            '@semantic-release/github',
            {
                assets: [
                    { path: 'CHANGELOG.md', label: 'Changelog' }
                ]
            }
            // GitHub releases will be published with the file CHANGELOG.md
        ]
    ]
}
