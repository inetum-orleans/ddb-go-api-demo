local ddb = import 'ddb.docker.libjsonnet';

local domain = std.join('.', [std.extVar('core.domain.sub'), std.extVar('core.domain.ext')]);

local app_workdir = '/app';

ddb.Compose() {
  services: {
    db: ddb.Build('db') + ddb.User()
        + ddb.Binary('mysql', app_workdir, 'mysql  -hdb -u' + std.extVar('app.db.user') + ' -p' + std.extVar('app.db.password'))
        + ddb.Binary('mysqldump', app_workdir, 'MYSQL_PWD=' + std.extVar('app.db.password') + ' mysqldump  -hdb -u' + std.extVar('app.db.user') + ' ' + std.extVar('app.db.user'))
        + ddb.Expose('3306')
        + {
          environment: {
            MYSQL_ROOT_PASSWORD: std.extVar('app.db.password'),
            MYSQL_DATABASE: std.extVar('app.db.name'),
            MYSQL_USER: std.extVar('app.db.user'),
            MYSQL_PASSWORD: std.extVar('app.db.password'),
          },
          volumes: [
            'db-data:/var/lib/mysql:rw',
            ddb.path.project + ':' + app_workdir,
          ],
        },
    go: ddb.Build('go')
        + ddb.User()
        + ddb.Binary('go', '/var/www/html', 'go')
        + ddb.Binary('air', '/var/www/html', 'air', exe=true)
        + ddb.Expose(std.extVar('app.api.port') + ':' + std.extVar('app.api.port'))
        + {
          volumes+: [
            ddb.path.project + ':/var/www/html',
          ],
        },
  },
}
