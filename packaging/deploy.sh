#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

CURDIR=`pwd`
GOIARDI_VERSION=`git describe --long --always`

gem install package_cloud

if [ -z ${PACKAGECLOUD_REPO} ] ; then
  echo "The environment variable PACKAGECLOUD_REPO must be set."
  exit 1
fi

# debian/raspbian
package_cloud push ${PACKAGECLOUD_REPO}/debian/wheezy ${DIR}/artifacts/wheezy/*.deb

# debian/jessie
package_cloud push ${PACKAGECLOUD_REPO}/debian/jessie ${DIR}/artifacts/jessie/*.deb
package_cloud push ${PACKAGECLOUD_REPO}/ubuntu/xenial ${DIR}/artifacts/jessie/*amd64.deb

# ubuntu
package_cloud push ${PACKAGECLOUD_REPO}/ubuntu/trusty ${DIR}/artifacts/trusty/*.deb

# centos

package_cloud push ${PACKAGECLOUD_REPO}/el/6 ${DIR}/artifacts/el6/*.rpm
package_cloud push ${PACKAGECLOUD_REPO}/el/7 ${DIR}/artifacts/el7/*.rpm
