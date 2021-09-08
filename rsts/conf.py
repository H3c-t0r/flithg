# -*- coding: utf-8 -*-
#
# Configuration file for the Sphinx documentation builder.
#
# This file does only contain a selection of the most common options. For a
# full list see the documentation:
# http://www.sphinx-doc.org/en/stable/config

# -- Path setup --------------------------------------------------------------

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use os.path.abspath to make it absolute, like shown here.
#
# import os
# import sys
# sys.path.insert(0, os.path.abspath('.'))

import sphinx.application
import sphinx.errors

sphinx.application.ExtensionError = sphinx.errors.ExtensionError

# -- Project information -----------------------------------------------------

project = u"Flyte"
copyright = u"2021, Flyte Authors"
author = u"Flyte"

# The short X.Y version
version = u""
# The full version, including alpha/beta/rc tags
release = u"0.16.0"

# -- General configuration ---------------------------------------------------

# If your documentation needs a minimal Sphinx version, state it here.
#
# needs_sphinx = '1.0'

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
    "sphinx.ext.autodoc",
    "sphinx.ext.autosectionlabel",
    "sphinx.ext.doctest",
    "sphinx.ext.intersphinx",
    "sphinx.ext.todo",
    "sphinx.ext.coverage",
    "sphinx.ext.ifconfig",
    "sphinx.ext.viewcode",
    "sphinx.ext.extlinks",
    "sphinx.ext.napoleon",
    "sphinx.ext.doctest",
    "sphinx.ext.intersphinx",
    "sphinx.ext.coverage",
    "sphinx-prompt",
    "sphinx_copybutton",
    "sphinx_tabs.tabs",
    "sphinxext.remoteliteralinclude",
    "sphinx_issues",
    # "sphinx_search.extension",
    "sphinx_panels",
    # "sphinxcontrib.mermaid",
    "sphinxcontrib.yt",
]

extlinks = {
    "propeller": ("https://github.com/flyteorg/flytepropeller/tree/master/%s", ""),
    "stdlib": ("https://github.com/flyteorg/flytestdlib/tree/master/%s", ""),
    "kit": ("https://github.com/flyteorg/flytekit/tree/master/%s", ""),
    "plugins": ("https://github.com/flyteorg/flyteplugins/tree/v0.1.4/%s", ""),
    "idl": ("https://github.com/flyteorg/flyteidl/tree/v0.14.1/%s", ""),
    "admin": ("https://github.com/flyteorg/flyteadmin/tree/master/%s", ""),
}

# Add any paths that contain templates here, relative to this directory.
# templates_path = ['_templates']

# The suffix(es) of source filenames.
# You can specify multiple suffix as a list of string:
#
# source_suffix = ['.rst', '.md']
source_suffix = ".rst"

# The master toctree document.
master_doc = "index"

# The language for content autogenerated by Sphinx. Refer to documentation
# for a list of supported languages.
#
# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
language = None

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path .
exclude_patterns = [u"_build", "Thumbs.db", ".DS_Store"]

# -- Options for HTML output -------------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_favicon = "images/flyte_circle_gradient_1_4x4.png"
html_logo = "images/flyte_circle_gradient_1_4x4.png"
html_theme = "furo"
html_title = "Flyte Docs"

templates_path = ["_templates"]

pygments_style = "tango"
pygments_dark_style = "native"

html_theme_options = {
    "light_css_variables": {
        "color-brand-primary": "#4300c9",
        "color-brand-content": "#4300c9",
    },
    "dark_css_variables": {
        "color-brand-primary": "#9D68E4",
        "color-brand-content": "#9D68E4",
    },
    # custom flyteorg furo theme options
    "github_repo": "flytesnacks",
    "github_username": "flyteorg",
    "github_commit": "master",
    "docs_path": "cookbook/docs",  # path to documentation source
    "sphinx_gallery_src_dir": "cookbook",  # path to directory of sphinx gallery source files relative to repo root
    "sphinx_gallery_dest_dir": "auto",  # path to root directory containing auto-generated sphinx gallery examples
}

# Theme options are theme-specific and customize the look and feel of a theme
# further.  For a list of options available for each theme, see the
# documentation.
#
# html_theme_options = {}

# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
html_static_path = ["_static"]
html_css_files = ["custom.css"]

# Custom sidebar templates, must be a dictionary that maps document names
# to template names.
#
# The default sidebars (for documents that don't match any pattern) are
# defined by theme itself.  Builtin themes are using these templates by
# default: ``['localtoc.html', 'relations.html', 'sourcelink.html',
# 'searchbox.html']``.
#
# html_sidebars = {"**": ["logo-text.html", "globaltoc.html", "localtoc.html", "searchbox.html"]}

# -- Options for HTMLHelp output ---------------------------------------------

# Output file base name for HTML help builder.
htmlhelp_basename = "Flytedoc"

# -- Options for LaTeX output ------------------------------------------------

latex_elements = {
    # The paper size ('letterpaper' or 'a4paper').
    #
    # 'papersize': 'letterpaper',
    # The font size ('10pt', '11pt' or '12pt').
    #
    # 'pointsize': '10pt',
    # Additional stuff for the LaTeX preamble.
    #
    # 'preamble': '',
    # Latex figure (float) alignment
    #
    # 'figure_align': 'htbp',
}

# Grouping the document tree into LaTeX files. List of tuples
# (source start file, target name, title,
#  author, documentclass [howto, manual, or own class]).
latex_documents = [
    (master_doc, "Flyte.tex", u"Flyte Documentation", u"Flyte Authors", "manual"),
]

# -- Options for manual page output ------------------------------------------

# One entry per manual page. List of tuples
# (source start file, name, description, authors, manual section).
man_pages = [(master_doc, "flyte", u"Flyte Documentation", [author], 1)]

# -- Options for Texinfo output ----------------------------------------------

# Grouping the document tree into Texinfo files. List of tuples
# (source start file, target name, title, author,
#  dir menu entry, description, category)
texinfo_documents = [
    (
        master_doc,
        "Flyte",
        u"Flyte Documentation",
        author,
        "Flyte",
        "Accelerate your ML and data workflows to production.",
        "Miscellaneous",
    ),
]

# -- Extension configuration -------------------------------------------------
autosectionlabel_prefix_document = True
autosectionlabel_maxdepth = 2

# -- Options for intersphinx extension ---------------------------------------

# Example configuration for intersphinx: refer to the Python standard library.
# intersphinx configuration. Uncomment the repeats with the local paths and update your username
# to help with local development.
intersphinx_mapping = {
    "python": ("https://docs.python.org/3", None),
    "numpy": ("https://numpy.org/doc/stable", None),
    "pandas": ("https://pandas.pydata.org/pandas-docs/stable/", None),
    "torch": ("https://pytorch.org/docs/master/", None),
    "scipy": ("https://docs.scipy.org/doc/scipy/reference", None),
    "matplotlib": ("https://matplotlib.org", None),
    "flytekit": ("https://flyte.readthedocs.io/projects/flytekit/en/latest/", None),
    # "flytekit": ("/Users/yourusername/go/src/github.com/flyteorg/flytekit/docs.rst/build/html", None),
    "flytectl": ("https://flytectl.readthedocs.io/en/latest/", None),
    # "flytectl": ("/Users/yourusername/go/src/github.com/flyteorg/flytectl/docs/build/html", None),
    "cookbook": ("https://flytecookbook.readthedocs.io/en/latest/", None),
    "flyteidl": ("https://docs.flyte.org/projects/flyteidl/en/latest", None),
}

# Sphinx-tabs config
sphinx_tabs_valid_builders = ["linkcheck"]

# Sphinx-mermaid config
# mermaid_output_format = 'raw'
# mermaid_version = 'latest'
# mermaid_init_js = "mermaid.initialize({startOnLoad:false});"

# Makes it so that only the command is copied, not the output
copybutton_prompt_text = "$ "

# -- Options for todo extension ----------------------------------------------

# If true, `todo` and `todoList` produce output, else they produce nothing.
todo_include_todos = True

# -- Options for Sphix issues extension --------------------------------------

# GitHub repo
issues_github_path = "flyteorg/flyte"

# equivalent to
issues_uri = "https://github.com/flyteorg/flyte/issues/{issue}"
issues_pr_uri = "https://github.com/flyteorg/flyte/pull/{pr}"
issues_commit_uri = "https://github.com/flyteorg/flyte/commit/{commit}"

# Usage:
# See issue :issue:`42`
# See issues :issue:`12,13`
# See :issue:`sloria/konch#45`.
# See PR :pr:`58`
