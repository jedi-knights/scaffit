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
        '@semantic-release/github',
    ],

    // Responsible for determining the type of release to make (major, minor or patch)
    // See: https://github.com/semantic-release/commit-analyzer#configuration
    analyzeCommits: [
        '@semantic-release/commit-analyzer',
    ],

    // Responsible for generating the content of the release note
    generateNotes: [
        '@semantic-release/release-notes-generator',
    ],

    // Responsible for preparing the release, for example creating or updating files such as package.json, CHANGELOG.md, documentation or compiled assets
    prepare: [
        '@semantic-release/changelog',
        '@semantic-release/git',
    ],

    // Responsible for preparing the release, for example creating or updating files such as package.json, CHANGELOG.md, documentation or compiled assets
    publish: [
        '@semantic-release/github',
    ],

    plugins: [
        '@semantic-release/commit-analyzer',
        '@semantic-release/release-notes-generator',
        '@semantic-release/changelog',
        [
            '@semantic-release/exec',
            {
                prepareCmd: 'echo ${nextRelease.version} > VERSION',
            },
        ],
        '@semantic-release/npm',
        ['@semantic-release/git', {
            assets: ['./public/**/*', 'package.json', 'package-lock.json', 'CHANGELOG.md', 'VERSION'],
            message: 'chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}',
        }],
        '@semantic-release/github',
    ],
};