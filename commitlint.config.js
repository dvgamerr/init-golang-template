// Commitlint configuration with JIRA ID support
// https://commitlint.js.org/
// Format: JIRA-123/type: subject
// Example: JIRA-456/fix: resolve authentication issue

module.exports = {
  extends: ['@commitlint/config-conventional'],
  parserPreset: {
    parserOpts: {
      // Custom parser to handle JIRA-ID/type: format
      headerPattern: /^(?:([A-Z]+-\d+)\/)?([\w]+)(?:\(([^\)]*)\))?:\s(.*)$/,
      headerCorrespondence: ['jira', 'type', 'scope', 'subject'],
    },
  },
  rules: {
    // JIRA ID rules
    'jira-id-format': [2, 'always'],

    // Type rules
    'type-enum': [
      2,
      'always',
      [
        'feat',     // New feature
        'fix',      // Bug fix
        'docs',     // Documentation changes
        'style',    // Code style changes (formatting, missing semi colons, etc)
        'refactor', // Code refactoring
        'perf',     // Performance improvements
        'test',     // Adding or updating tests
        'build',    // Build system or external dependencies
        'ci',       // CI configuration files and scripts
        'chore',    // Other changes that don't modify src or test files
        'revert',   // Revert a previous commit
      ],
    ],
    'type-case': [2, 'always', 'lower-case'],
    'type-empty': [2, 'never'],

    // Scope rules
    'scope-case': [2, 'always', 'lower-case'],

    // Subject rules
    'subject-empty': [2, 'never'],
    'subject-full-stop': [2, 'never', '.'],
    'subject-case': [0], // Disable subject case checking

    // Header rules
    'header-max-length': [2, 'always', 120],

    // Body rules
    'body-leading-blank': [1, 'always'],
    'body-max-line-length': [2, 'always', 100],

    // Footer rules
    'footer-leading-blank': [1, 'always'],
    'footer-max-line-length': [2, 'always', 100],
  },
  plugins: [
    {
      rules: {
        'jira-id-format': (parsed) => {
          const { jira, raw } = parsed;

          // If JIRA ID is present, validate format
          if (jira) {
            // JIRA ID should match pattern: PROJECT-123
            const jiraRegex = /^[A-Z]+-\d+$/;
            if (!jiraRegex.test(jira)) {
              return [
                false,
                `JIRA ID "${jira}" must match format: PROJECT-123 (uppercase letters, hyphen, numbers)`,
              ];
            }
          }

          // JIRA ID is optional, so it's valid even if not present
          return [true];
        },
      },
    },
  ],
};
