#!/bin/bash
#
# same result as 'npm run build-mv' under services/ui
#

function usage() {
    cat <<EOF

Usage: $0 [-m] [-h] [dev|test|prod|clear|help]

No argument runs 'vite build'.  With argument runs 'vite build --mode arg',
where arg is dev, test, or prod. All commands result in moving the resulting
'./dist' files into ./_media/ui/.

  -m:    build and move to ${MEDIA}
  -h:    display this help message
  dev:   vite build using mode dev using env vars from .env.dev
  test:  vite build using mode test using env vars from .env.test
  prod:  vite build using mode prod using env vars from .env.prod
  clear: remove all files from '${MEDIA}'

  There is a local variable PREFIX (${PREFIX}) that references the relative
  path to package.json. There is another local variable MEDIA (${MEDIA}) that
  references the relative path to media files.

  These two local variables should be modified to fit your repo's needs.

EOF
    exit 0
}


PREFIX=./services/ui
MEDIA=./_media/shared/ui

MOVE=0

# Parse optional flags first
while getopts "mh" opt; do
  case $opt in
    m)
        MOVE=1
        ;;
    h)
        usage
        ;;
    *)
        usage
        ;;
  esac
done
shift $((OPTIND - 1))


if [ $# -eq 0 ]; then
    if [ ${MOVE} -eq 1 ]; then
        npm run --prefix "${PREFIX}" build-mv
    else
        npm run --prefix "${PREFIX}" build
    fi
    exit 0
fi

for cmd in "$@"; do
    case $cmd in
    dev)
        if [ ${MOVE} -eq 1 ]; then
            npm run --prefix ${PREFIX} build-mv-dev
        else
            npm run --prefix ${PREFIX} build-dev
        fi
        exit 0
        ;;
    test)
        if [ ${MOVE} -eq 1 ]; then
            npm run --prefix ${PREFIX} build-mv-test
        else
            npm run --prefix ${PREFIX} build-test
        fi
        exit 0
        ;;
    prod)
        if [ ${MOVE} -eq 1 ]; then
            npm run --prefix ${PREFIX} build-mv-prod
        else
            npm run --prefix ${PREFIX} build-prod
        fi
        exit 0
        ;;
    clear)
        find ${MEDIA} -mindepth 1 ! -name '.gitkeep' -exec rm -rf {} +
        exit 0
        ;;
    esac
done
