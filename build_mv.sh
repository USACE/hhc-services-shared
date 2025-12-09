#!/bin/bash
#
# same result as 'npm run build-mv' under services/ui
#

function usage() {
    cat <<EOF

Usage: $0 [dev|test|prod|clear|help]

No argument runs 'vite build'.  With argument runs 'vite build --mode arg',
where arg is dev, test, or prod. All commands result in moving the resulting
'./dist' files into ./_media/ui/.

  dev:   vite build using mode dev using env vars from .env.dev
  test:  vite build using mode test using env vars from .env.test
  prod:  vite build using mode prod using env vars from .env.prod
  clear: remove all files from '${MEDIA}'
  help:  display this help message

  There is a local variable PREFIX (${PREFIX}) that references the relative
  path to package.json. There is another local variable MEDIA (${MEDIA}) that
  references the relative path to media files.

  These two local variables should be modified to fit your repo's needs.

EOF
    exit 0
}


PREFIX=./services/ui
MEDIA=./_media/shared/ui

if [ $# -eq 0 ]; then
    npm run --prefix ${PREFIX} build-mv
    exit 0
fi

for cmd in "$@"; do
    case $cmd in
    dev)
        npm run --prefix ${PREFIX} build-mv-dev
        exit 0
        ;;
    test)
        npm run --prefix ${PREFIX} build-mv-test
        exit 0
        ;;
    prod)
        npm run --prefix ${PREFIX} build-mv-prod
        exit 0
        ;;
    clear)
        find ${MEDIA} -mindepth 1 ! -name '.gitkeep' -exec rm -rf {} +
        exit 0
        ;;
    help)
        usage
        ;;
    esac
done
