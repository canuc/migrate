set -e

ARCH=`uname -m`
go build -o terraform-provider-migrate

mkdir -p ~/.terraform.d/plugins/github.com/canuc/migrate/0.1.0/darwin_${ARCH}/
mkdir -p ~/.terraform.d/providers/github.com/canuc/migrate/0.1.0/darwin_${ARCH}/
mkdir -p ~/.terraform.d/providers/github.com/canuc/migrate/0.1.0/darwin_${ARCH}/


cp terraform-provider-migrate ~/.terraform.d/plugins/github.com/canuc/migrate/0.1.0/darwin_${ARCH}/terraform-provider-migrate
cp terraform-provider-migrate ~/.terraform.d/providers/github.com/canuc/migrate/0.1.0/darwin_${ARCH}/terraform-provider-migrate

chmod a+x ~/.terraform.d/plugins/github.com/canuc/migrate/0.1.0/darwin_${ARCH}/terraform-provider-migrate
chmod a+x ~/.terrafgit 