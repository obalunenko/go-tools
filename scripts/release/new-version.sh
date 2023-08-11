#!/usr/bin/env bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

APP=${APP_NAME}

RELEASE_BRANCH=${RELEASE_BRANCH:-"release"}

echo "${SCRIPT_NAME} is running fo ${APP}... "

checkInstalled 'svu'

echo "${SCRIPT_NAME} is running... "

function requireReleaseBranch() {
  err=0
  branch=$(git branch --show-current)

  echo "Current branch is: ${branch}"

  if [[ ${branch} != "${RELEASE_BRANCH}" ]]; then
    err=1
  fi

   if [[ ${err} == 1 ]]; then
        echo >&2 "Please checkout to ${RELEASE_BRANCH} branch."
        exit 1
    fi
}

function require_clean_work_tree() {
  # Update the index
  git update-index -q --ignore-submodules --refresh
  err=0

  # Disallow unstagged changes in the working tree
  if ! git diff-files --quiet --ignore-submodules --; then
    echo >&2 "cannot $1: you have unstaged changes."
    git diff-files --name-status -r --ignore-submodules -- >&2
    err=1
  fi

  # Disallow uncommitted changes in the index
  if ! git diff-index --cached --quiet HEAD --ignore-submodules --; then
    echo >&2 "cannot $1: your index contains uncommitted changes."
    git diff-index --cached --name-status -r --ignore-submodules HEAD -- >&2
    err=1
  fi

  if [[ ${err} == 1 ]]; then
    echo >&2 "Please commit or stash them."
    exit 1
  fi
}

function menu() {
  PREV_VERSION=$(svu current)
  clear

  echo "Current version: ${PREV_VERSION}"
  printf "Select what you want to update: \n"
  printf "1 - Major update\n"
  printf "2 - Minor update\n"
  printf "3 - Patch update\n"
  printf "4 - Exit\n"
  read -r selection

  case "$selection" in
  1)
    printf "Major updates......\n"
    NEW_VERSION=$(svu major)
    ;;
  2)
    printf "Run Minor update.........\n"
    NEW_VERSION=$(svu minor)
    ;;
  3)
    printf "Patch update.........\n"
    NEW_VERSION=$(svu patch)
    ;;
  4)
    printf "Exit................................\n"
    exit 1
    ;;
  *)
    clear
    printf "Incorrect selection. Try again\n"
    menu
    ;;
  esac

}

## Check if release branch
requireReleaseBranch

## Check if git is clean
require_clean_work_tree "create new version"

git pull

## Sem ver update menu
menu

NEW_TAG=${NEW_VERSION}

echo "New version is: ${NEW_TAG}"
while true; do
  echo "Is it ok? (:y)?:"
  read -r yn
  case $yn in
  [Yy]*)

  git tag -a "${NEW_TAG}" -m "${NEW_TAG}" &&
      git push --tags

    break
    ;;
  [Nn]*)
    echo "Cancel"
    break
    ;;
  *)
    echo "Please answer yes or no."
    ;;
  esac
done

echo "${SCRIPT_NAME} done."
