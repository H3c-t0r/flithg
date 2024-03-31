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

import os
import logging
import sys

import sphinx.application
import sphinx.errors
from sphinx.util import logging as sphinx_logging


sys.path.insert(0, os.path.abspath("../"))
sys.path.append(os.path.abspath("./_ext"))

sphinx.application.ExtensionError = sphinx.errors.ExtensionError

# -- Project information -----------------------------------------------------

project = "Flyte"
copyright = "2022, Flyte Authors"
author = "Flyte"

# The short X.Y version
version = ""
# The full version, including alpha/beta/rc tags
release = "1.11.1-b0"

# -- General configuration ---------------------------------------------------

# If your documentation needs a minimal Sphinx version, state it here.
#
# needs_sphinx = '1.0'

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
    "sphinx.ext.napoleon",
    "sphinx.ext.autodoc",
    "sphinx.ext.autosummary",
    "sphinx.ext.autosectionlabel",
    "sphinx.ext.doctest",
    "sphinx.ext.inheritance_diagram",
    "sphinx.ext.intersphinx",
    "sphinx.ext.graphviz",
    "sphinx.ext.todo",
    "sphinx.ext.coverage",
    "sphinx.ext.ifconfig",
    "sphinx.ext.viewcode",
    "sphinx.ext.extlinks",
    "sphinx-prompt",
    "sphinx_copybutton",
    "sphinxext.remoteliteralinclude",
    "sphinx_issues",
    "sphinx_click",
    "sphinx_panels",
    "sphinxcontrib.mermaid",
    "sphinxcontrib.video",
    "sphinxcontrib.youtube",
    "sphinx_tabs.tabs",
    "sphinx_tags",
    "myst_nb",
    # custom extensions
    "auto_examples",
    "import_projects",
]

source_suffix = {
    ".rst": "restructuredtext",
    ".md": "myst-nb",
}

extlinks = {
    "propeller": ("https://github.com/flyteorg/flytepropeller/tree/master/%s", ""),
    "stdlib": ("https://github.com/flyteorg/flytestdlib/tree/master/%s", ""),
    "kit": ("https://github.com/flyteorg/flytekit/tree/master/%s", ""),
    "plugins": ("https://github.com/flyteorg/flyteplugins/tree/v0.1.4/%s", ""),
    "idl": ("https://github.com/flyteorg/flyteidl/tree/v0.14.1/%s", ""),
    "admin": ("https://github.com/flyteorg/flyteadmin/tree/master/%s", ""),
    "cookbook": ("https://flytecookbook.readthedocs.io/en/latest/", None),
}


autosummary_generate = True
suppress_warnings = ["autosectionlabel.*", "myst.header"]
autodoc_typehints = "description"

# The master toctree document.
master_doc = "index"

# The language for content autogenerated by Sphinx. Refer to documentation
# for a list of supported languages.
#
# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
language = "en"

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path .
exclude_patterns = [
    "_build",
    "Thumbs.db",
    ".DS_Store",
    "auto/**/*.ipynb",
    "auto/**/*.py",
    "auto/**/*.md",
    "examples/**/*.ipynb",
    "examples/**/*.py",
    "jupyter_execute/**",
    "README.md",
    "_projects/**",
    "_src/**",
    "examples/**",
    "flytesnacks/index.md",
    "flytesnacks/bioinformatics_examples.md",
    "flytesnacks/feature_engineering.md",
    "flytesnacks/flyte_lab.md",
    "flytesnacks/ml_training.md",
    "flytesnacks/README.md",
    "flytekit/**/README.md",
    "flytekit/_templates/**",
    "flytectl/index.rst",
    "protos/boilerplate/**",
    "protos/tmp/**",
    "protos/gen/**",
    "protos/docs/**/index.rst",
    "protos/index.rst",
    "api/flytekit/_templates/**",
    "api/flytekit/index.rst",
    "reference/index.rst",
]

# -- Options for HTML output -------------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_favicon = "images/favicon-flyte-docs.png"
html_logo = "images/favicon-flyte-docs.png"
html_theme = "furo"
html_title = "Flyte"

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
    # "github_repo": "flyte",
    # "github_username": "flyteorg",
    # "github_commit": "master",
    # "docs_path": "rsts",  # path to documentation source
    # "sphinx_gallery_src_dir": "cookbook",  # path to directory of sphinx gallery source files relative to repo root
    # "sphinx_gallery_dest_dir": "auto",  # path to root directory containing auto-generated sphinx gallery examples
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
html_css_files = ["custom.css", "flyte.css", "algolia.css"]
html_js_files = ["custom.js"]

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
    (master_doc, "Flyte.tex", "Flyte Documentation", "Flyte Authors", "manual"),
]

# -- Options for manual page output ------------------------------------------

# One entry per manual page. List of tuples
# (source start file, name, description, authors, manual section).
man_pages = [(master_doc, "flyte", "Flyte Documentation", [author], 1)]

# -- Options for Texinfo output ----------------------------------------------

# Grouping the document tree into Texinfo files. List of tuples
# (source start file, target name, title, author,
#  dir menu entry, description, category)
texinfo_documents = [
    (
        master_doc,
        "Flyte",
        "Flyte Documentation",
        author,
        "Flyte",
        "Accelerate your ML and data workflows to production.",
        "Miscellaneous",
    ),
]

# -- Extension configuration -------------------------------------------------
# autosectionlabel_prefix_document = True
autosectionlabel_maxdepth = 2

# Tags config
tags_create_tags = True
tags_extension = ["md", "rst"]
tags_page_title = "Tag"
tags_overview_title = "Pages by Tags"

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
    "pandera": ("https://pandera.readthedocs.io/en/stable/", None),
}

myst_enable_extensions = ["colon_fence"]

# Sphinx-mermaid config
mermaid_output_format = "raw"
mermaid_version = "latest"
mermaid_init_js = "mermaid.initialize({startOnLoad:false});"

# Makes it so that only the command is copied, not the output
copybutton_prompt_text = "$ "

# prevent css style tags from being copied by the copy button
copybutton_exclude = 'style[type="text/css"]'

nb_execution_mode = "off"
nb_execution_excludepatterns = [
    "flytekit/**/*",
    "flytesnacks/**/*",
    "examples/**/*",
]
nb_custom_formats = {
    ".md": ["jupytext.reads", {"fmt": "md:myst"}],
}

# Pattern for removing intersphinx references from source files.
# This should handle cases like:
#
# - :ref:`cookbook:label` -> :ref:`label`
# - :ref:`Text <cookbook:label>` -> :ref:`Text <label>`
INTERSPHINX_REFS_PATTERN = r"([`<])(flyte:|flytekit:|flytectl:|flyteidl:|cookbook:|idl:)"
INTERSPHINX_REFS_REPLACE = r"\1"

# Pattern for replacing all ref/doc labels that point to protos/docs with /protos/docs
PROTO_REF_PATTERN = r"([:<])(protos/docs)"
PROTO_REF_REPLACE = r"\1/protos/docs"

# These patterns are used to replace values in source files that are imported
# from other repos.
REPLACE_PATTERNS = {

    r"<flyte:deployment/index>": r"</deployment/index>",
    r"<flytectl:index>": r"</flytectl/overview>",
    INTERSPHINX_REFS_PATTERN: INTERSPHINX_REFS_REPLACE,
    r"<auto_examples": r"<flytesnacks/examples",
    r"<protos/docs/core/core:taskmetadata>": r"<ref_flyteidl.core.TaskMetadata>",
    r"<protos/docs/core/core:tasktemplate>": r"<ref_flyteidl.core.TaskTemplate>",
    r"<flytesnacks/examples": r"</flytesnacks/examples",
    r"<auto_examples/basics/index>": r"</flytesnacks/examples/basics/index>",
    r"<deploy-sandbox-local>": r"<deployment-deployment-sandbox>",
    r"<deployment/configuration/general:configurable resource types>": r"<deployment-configuration-general>",
    r"<_tags/DistributedComputing>": r"</_tags/DistributedComputing>",
    r"{ref}`bioinformatics <bioinformatics>`": r"bioinformatics",
    PROTO_REF_PATTERN: PROTO_REF_REPLACE,
    r"/protos/docs/service/index": r"/protos/docs/service/service",
    r"<weather_forecasting>": r"</flytesnacks/weather_forecasting>",
}

# r"<environment_setup>": r"</flytesnacks/environment_setup>",

import_projects_config = {
    "clone_dir": "_projects",
    "flytekit_api_dir": "_src/flytekit/",
    "source_regex_mapping": REPLACE_PATTERNS,
    "list_table_toc": [
       "flytesnacks/tutorials",
       "flytesnacks/integrations",
       "flytesnacks/deprecated_integrations"
    ],
    "dev_build": bool(int(os.environ.get("MONODOCS_DEV_BUILD", 1))),
}

# Define these environment variables to use local copies of the projects. This
# is useful for building the docs in the CI/CD of the corresponding repos.
flytesnacks_local_path = os.environ.get("FLYTESNACKS_LOCAL_PATH", None)
flytekit_local_path = os.environ.get("FLYTEKIT_LOCAL_PATH", None)
flytectl_local_path = os.environ.get("FLYTECTL_LOCAL_PATH", None)

flytesnacks_path = flytesnacks_local_path or "_projects/flytesnacks"
flytekit_path = flytekit_local_path or "_projects/api/flytekit"

import_projects = [
    {
        "name": "flytesnacks",
        "source": flytesnacks_local_path or "https://github.com/flyteorg/flytesnacks",
        "docs_path": "docs",
        "dest": "flytesnacks",
        "cmd": [
            ["cp", "-R", f"{flytesnacks_path}/examples", "./examples"],
            [
                # remove un-needed docs files in flytesnacks
                "rm",
                "-rf",
                "flytesnacks/jupyter_execute",
                "flytesnacks/auto_examples",
                "flytesnacks/_build",
                "flytesnacks/_tags",
                "flytesnacks/getting_started",
                "flytesnacks/userguide.md",
                "flytesnacks/environment_setup.md",
                "flytesnacks/index.md",
                "examples/advanced_composition",
                "examples/basics",
                "examples/customizing_dependencies",
                "examples/data_types_and_io",
                "examples/development_lifecycle",
                "examples/extending",
                "examples/productionizing",
                "examples/testing",
                "flytesnacks/examples/advanced_composition",
                "flytesnacks/examples/basics",
                "flytesnacks/examples/customizing_dependencies",
                "flytesnacks/examples/data_types_and_io",
                "flytesnacks/examples/development_lifecycle",
                "flytesnacks/examples/extending",
                "flytesnacks/examples/productionizing",
                "flytesnacks/examples/testing",
            ]
        ],
        "local": flytesnacks_local_path is not None,
    },
    {
        "name": "flytekit",
        "source": flytekit_local_path or "https://github.com/flyteorg/flytekit",
        "docs_path": "docs/source",
        "dest": "api/flytekit",
        "cmd": [
            ["mkdir", "-p", import_projects_config["flytekit_api_dir"]],
            ["cp", "-R", f"{flytekit_path}/flytekit", import_projects_config["flytekit_api_dir"]],
            ["cp", "-R", f"{flytekit_path}/plugins", import_projects_config["flytekit_api_dir"]],
            ["cp", "-R", f"{flytekit_path}/tests", "./tests"],
        ],
        "local": flytekit_local_path is not None,
    },
    {
        "name": "flytectl",
        "source": flytectl_local_path or "https://github.com/flyteorg/flytectl",
        "docs_path": "docs/source",
        "dest": "flytectl",
        "local": flytectl_local_path is not None,
    },
    {
        "name": "flyteidl",
        "source": "../flyteidl",
        "docs_path": "protos",
        "dest": "protos",  # to stay compatible with flyteidl docs path naming
        "cmd": ["cp", "../flyteidl/README.md", "protos/README.md"],
        "local": True,
    }
]

# myst notebook docs customization
auto_examples_dir_root = "examples"
auto_examples_dir_dest = "flytesnacks/examples"
auto_examples_refresh = False

# -- Options for todo extension ----------------------------------------------

# If true, `todo` and `todoList` produce output, else they produce nothing.
todo_include_todos = True

# -- Options for Sphinx issues extension --------------------------------------

# GitHub repo
issues_github_path = "flyteorg/flyte"

# equivalent to
issues_uri = "https://github.com/flyteorg/flyte/issues/{issue}"
issues_pr_uri = "https://github.com/flyteorg/flyte/pull/{pr}"
issues_commit_uri = "https://github.com/flyteorg/flyte/commit/{commit}"


# Disable warnings from flytekit
os.environ["FLYTE_SDK_LOGGING_LEVEL_ROOT"] = "50"

# Disable warnings from tensorflow
os.environ["TF_CPP_MIN_LOG_LEVEL"] = "3"


class CustomWarningSuppressor(logging.Filter):
    """Filter logs by `suppress_warnings`."""

    def __init__(self, app: sphinx.application.Sphinx) -> None:
        self.app = app
        super().__init__()

    def filter(self, record: logging.LogRecord) -> bool:
        msg = record.getMessage()

        # TODO: These are all warnings that should be fixed as follow-ups to the
        # monodocs build project.
        filter_out = (
            "duplicate label",
            "Unexpected indentation",
            'Error with CSV data in "csv-table" directive',
            "Definition list ends without a blank line",
            "autodoc: failed to import module 'awssagemaker' from module 'flytekitplugins'",
            "Enumerated list ends without a blank line",
            'Unknown directive type "toc".',  # need to fix flytesnacks/contribute.md
        )

        if msg.strip().startswith(filter_out):
            return False

        if (
            msg.strip().startswith("document isn't included in any toctree")
            and record.location == "_tags/tagsindex"
        ):
            # ignore this warning, since we don't want the side nav to be
            # cluttered with the tags index page.
            return False

        return True


def setup(app: sphinx.application.Sphinx) -> None:
    """Setup root logger for Sphinx"""
    logger = logging.getLogger("sphinx")

    warning_handler, *_ = [
        h for h in logger.handlers
        if isinstance(h, sphinx_logging.WarningStreamHandler)
    ]
    warning_handler.filters.insert(0, CustomWarningSuppressor(app))
