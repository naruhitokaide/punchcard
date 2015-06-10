
import os
import pytest

from .git import git


# TODO do not use cwd, use random subdir
@pytest.fixture
def repo():
    return git(os.path.abspath(os.getcwd()))


def test_creation(repo):
    assert repo is not None


# def test_init(repo):
#     repo.init()
#     assert os.path.exists(os.path.join(os.path.abspath(os.getcwd()), '.git'))


def test_add(repo):
    pass


def test_commit(repo):
    pass
