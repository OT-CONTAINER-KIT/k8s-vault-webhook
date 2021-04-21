# Configuration file for the Sphinx documentation builder.
#
# This file only contains a selection of the most common options. For a full
# list see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Path setup --------------------------------------------------------------

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use os.path.abspath to make it absolute, like shown here.
#

import sphinx_rtd_theme

project = 'K8s Vault Webhook'
copyright = '2021, Opstree Solutions'
author = 'Opstree Solutions'

release = 'stable'

source_suffix = ['.rst', '.md']
templates_path = ['_templates']
html_static_path = ["_static"]


master_doc = 'index'

exclude_patterns = []

pygments_style = 'sphinx'

htmlhelp_basename = 'K8sVaultOperator'
extensions = [
    'sphinx.ext.autodoc',
    'sphinx.ext.doctest',
    'sphinx.ext.intersphinx',
    'sphinx.ext.todo',
    'sphinx.ext.coverage',
    'sphinx.ext.mathjax',
    'sphinx.ext.ifconfig',
    'sphinx.ext.viewcode',
    'sphinx.ext.githubpages',
]

html_theme = "sphinx_rtd_theme"
html_theme_path = [sphinx_rtd_theme.get_html_theme_path()]
html_theme_options = {
    'analytics_anonymize_ip': False,
    'navigation_depth': 5,
    'collapse_navigation': False,
    'sticky_navigation': True,
}
html_context = {}
html_static_path = ['_static']


def setup(app):
    app.add_css_file('width.css')
