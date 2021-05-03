const { description } = require('../../package')

module.exports = {
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#title
   */
  title: 'K8s Vault Webhook',
  base: '/k8s-vault-webhook/',
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#description
   */
  description: description,

  /**
   * Extra tags to be injected to the page HTML `<head>`
   *
   * ref：https://v1.vuepress.vuejs.org/config/#head
   */
  head: [
    ['meta', { name: 'theme-color', content: '#0065b3' }],
    ['meta', { name: 'apple-mobile-web-app-capable', content: 'yes' }],
    ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }]
  ],

  /**
   * Theme configuration, here is the default theme configuration for VuePress.
   *
   * ref：https://v1.vuepress.vuejs.org/theme/default-theme-config.html
   */
  themeConfig: {
    repo: 'https://github.com/OT-CONTAINER-KIT/k8s-vault-webhook',
    editLinks: false,
    docsDir: '',
    editLinkText: '',
    lastUpdated: false,
    nav: [
      {
        text: 'Guide',
        link: '/guide/',
      },
      {
        text: 'OpsTree',
        link: 'https://opstree.com'
      }
    ],
    sidebar: {
      '/guide/': [
        {
          title: 'Guide',
          collapsable: false,
          children: [
            '',
            'secret-manager',
            // 'vault-config',
            // 'webhook-install',
          ]
        },
        {
          title: 'Getting Started',
          collapsable: false,
          children: [
            'installation',
            'configuration',
            'annotations',
          ]
        },
        {
          title: 'Integration',
          collapsable: false,
          children: [
            'hashicorp-vault',
          ]
        },
        {
          title: 'Examples',
          collapsable: false,
          children: [
            'hashicorp-vault-example',
          ]
        },
        {
          title: 'Development',
          collapsable: false,
          children: [
            'development',
          ]
        },
        {
          title: 'Changelog',
          collapsable: false,
          children: [
            'changelog',
          ]
        }
      ],
    }
  },

  /**
   * Apply plugins，ref：https://v1.vuepress.vuejs.org/zh/plugin/
   */
  plugins: [
    '@vuepress/plugin-back-to-top',
    '@vuepress/plugin-medium-zoom',
  ]
}
