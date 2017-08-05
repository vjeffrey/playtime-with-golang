pkg_name=vj-text-adventure
pkg_description="silly game from tutorial"
pkg_origin=vjeffrey
pkg_version="0.0.0"
pkg_maintainer="Victoria Jeffrey <vjeffrey@chef.io>"
pkg_license=('UNLICENSED')
pkg_upstream_url="https://github.com/vjeffrey/playtime-with-golang"

pkg_interpreters=(bin/bash)
pkg_deps=(core/coreutils)
pkg_build_deps=(core/go/1.8 core/git core/make core/which)
pkg_bin_dirs=(bin)

do_begin() {
  export GOPATH="${HAB_CACHE_ARTIFACT_PATH}/.go"
  export GOBIN="${GOPATH}/bin"
  export REPO="${GOPATH}/src/github.com/playtime-with-golang"
}

do_download() {
  mkdir -p ${GOPATH}
  rm -rf ${REPO}
  git clone https://github.com/vjeffrey/playtime-with-golang.git ${REPO}
}

do_build() {
  pushd ${REPO}
  PATH="${PATH}:${REPO}" go build .
  popd
}

do_install() {
  pushd ${REPO}
  PATH="${PATH}:${REPO}" go get github.com/Sirupsen/logrus && go get github.com/lib/pq
  popd
}