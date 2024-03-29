const branch = process.env.TRAVIS_BRANCH

const config = {
  branches: ['main', { name: 'develop', prerelease: 'rfv' }, '+([0-9])?(.{+([0-9]),x}).x'],
  plugins: [['@semantic-release/commit-analyzer',
  {
    releaseRules: [
      {
        type: "refactor",
        release: "patch",
      },
    ],
  },
], 
  ['@semantic-release/release-notes-generator',
  {
    presetConfig: {
      types: [
        {
          type: "refactor",
          section: "Refactors",
          hidden: false,
        },
      ],
    },
  }],
  [
    "@semantic-release/exec",
    {
      prepareCmd: "printf '${nextRelease.version}' > VERSION",
      successCmd: "TRAVIS_TAG=${nextRelease.gitTag}",
    },
  ],
]}

if (branch == 'main') {
  config.plugins.push(['@semantic-release/changelog',
    {
      changelogFile: "CHANGELOG.md",
    },
  ],
  ['@semantic-release/git',
    {
      assets: ['CHANGELOG.md', 'VERSION'],
    },
  ])
}

// config.plugins.push(["@codedependant/semantic-release-docker",
//     {
//       dockerTags: ["{{git_tag}}", "main"],
//       dockerImage: "735283066345.dkr.ecr.us-east-2.amazonaws.com/spandigital/presidium-enterprise",
//       dockerFile: "build/docker/presidium.Dockerfile",
//     },
//   ],
//   ["@codedependant/semantic-release-docker",
//     {
//       dockerTags: ["{{git_tag}}", "main"],
//       dockerImage: "735283066345.dkr.ecr.us-east-2.amazonaws.com/spandigital/presidium-migrate",
//       dockerFile: "build/docker/migrate.Dockerfile",
//     },
//   ],
//   // ["semantic-release-jira-releases",
//   //   {
//   //     projectId: "PRSDM",
//   //     jiraHost: "spandigital.atlassian.net",
//   //     ticketPrefixes: ["PRSDM"],
//   //     releaseNameTemplate: "v${version}",
//   //     releaseDescriptionTemplate: "Automated release with Semantic Release",
//   //   },
//   // ],
// )

module.exports = config
