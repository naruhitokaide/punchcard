
__author__ = 'Christian Rotzoll'
__version__ = '0.1.0'

import os
import arrow
import click
import subprocess

from git import git


@click.group()
@click.version_option(version=__version__)
def cli():
    pass


@cli.command()
@click.option('--min', default=0, help='Smallest number of commits per day.')
@click.option('--max', default=10, help='Largest number of commits per day.')
@click.option('--location', default='.', type=click.Path(),
    help='Location of the git project, which will be created')
def random(min, max, location):
    # find better names for min and max to avoid shadowing built in functions
    click.echo('random: ' + ' '.join([str(min), str(max), location]))
    # TODO use either relative or absolute path, but pass only abs path
    random_git = git(location)
    random_git.init()


if __name__ == '__main__':
    cli()
