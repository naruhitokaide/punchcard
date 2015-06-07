
__author__ = 'Christian Rotzoll'
__version__ = '0.1.0'

import os
import arrow
import click
import subprocess


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
    click.echo('random: ' + ' '.join([str(min), str(max), location]))
    git_init(location)


def git_init(path):
    assert os.path.isabs(path)
    click.echo('git init: ' + path)
    with cd(path):
        subprocess.call('git init')


def git_add(filenames):
    for filename in filenames:
        assert os.path.isfile(filename)  # needs entire path
        subprocess.call('git add ' + filename)

def git_commit(date):
    click.echo('git commit: ' + date)
    subprocess.call("git commit -m '' ")


class cd:
    """Context manager for changing the current working directory"""
    def __init__(self, newPath):
        self.newPath = os.path.expanduser(newPath)

    def __enter__(self):
        self.savedPath = os.getcwd()
        os.chdir(self.newPath)

    def __exit__(self, etype, value, traceback):
        os.chdir(self.savedPath)


if __name__ == '__main__':
    cli()
