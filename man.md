zabbixctl(1) -- tool for working with zabbix using command line interface
===========

## DESCRIPTIOPN

*zabbixctl* is a tool for working with zabbix server api using command line
interface, it provides effective way for operating on statuses of triggers,
hosts latest data and groups of users.

## CONFIGURATION

*zabbixctl* reads its configuration from  ~/.config/zabbixctl.conf and must be
written using following syntax:

    [server]
      address  = "https://zabbix.local"
      username = "admin"
      password = "password"

    [session]
      path = "~/.cache/zabbixctl.session"

  *zabbixctl* will authorize against 'zabbix.local' using given user
credentials and save a Zabbix session to ~/.cache/zabbixctl.session.
Subsequent runs will use saved session instead of new re-authorizing.
Zabbix sessions have a default TTL of 15 minutes, so if the saved Zabbix
session is outdated, zabbixctl will repeat authorization and rewrite the
session file.

## SYNOPSIS

    zabbixctl [options] -T [/<pattern>...]
    zabbixctl [options] -L <hostname>... [/<pattern>...]
    zabbixctl [options] -G [/<pattern>...]
    zabbixctl -h | --help
    zabbixctl --version

## OPTIONS

**-T --triggers**

Search Zabbix triggers statuses. Triggers can be filtered usingi the /<pattern>
argument, for example, search and acknowledge all triggers in a problem state
and match the word 'cache':

`zabbixctl -Tp /cache`

**-y --only-nack**

Show only not acknowledged triggers.

**-x --severity**

Specify minimum trigger severity.  Once for information, twice for warning,
three for disaster, four for high, five for disaster.

**-p --problem**

Show triggers that have a problem state.

**-r --recent**

Show triggers that have recently been in a problem state.

**-s --since <date>**

Show triggers that have changed their state after the given time, default: 7
days ago.

**-u --until <date>**

Show triggers that have changed their state before the given time.

**-m --maintenance**

Show hosts in maintenance.

**-i --sort <fields>**

Show triggers sorted by specified fields, default: lastchange,priority.

**-o --order <order>**

Show triggers in specified order, default: DESC.

**-n --limit <amount>**

Show specified amount of triggers.

**-k --acknowledge**

Acknowledge all retrieved triggers.

**-f --noconfirm**

Do not prompt for acknowledge confirmation.

**-L --latest-data**

Search and show latest data for specified host(s). Hosts can be searched
using a wildcard character '*'.  Data can be filtered using the /<pattern>
argument, for example retrieve latest data for database nodes and search
information about replication:

`zabbixctl -L dbnode-* /replication`

**-g --graph**

Show links on graph pages.

**-G --groups**

Search and operate on configuration of usergroups.

**-l --list**

Show list of users in specified usergroup.

**-a --add**

Add specified <user> to specified usergroup.

**-r --remove**

Remove specified <user> from specified usergroup.

**-f --noconfirm**

Do not prompt for confirmation.

## COMMON OPTIONS

**-c --config <path>**

Use specified configuration file, default: `$HOME/.config/zabbixctl.conf`

**-v --verbosity**

Specify program output verbosity. Once for debug, twice for trace.

**-h --help**

Show this screen.

**--version**

Show version.

## EXAMPLES

*Listing triggers in a problem state*

```
zabbixctl -Tp
```

*Listing triggers that have recenty been in a problem state*

```
zabbixctl -Tr
```

*Listing and filtering triggers that contain the word mysql*

```
zabbixctl -T /mysql
```

*Listing and acknowledging triggers that severity level is DISASTER*

```
zabbixctl -T -xxxxx -k
```

*Listing latest data for db nodes and filtering for information about replication lag*

```
zabbixctl -L dbnode* /lag
```

*Opening stacked graph for CPU quote use of selected containers*

```
zabbixctl -L 'container-*' /cpu quota --stacked
```

*Listing usergroups that starts with 'HTTP_'*

```
zabbixctl -G HTTP_*
```

*Listing usergroups that contain user admin*

```
zabbixctl -G /admin
```

*Adding user admin to groups that contain user guest*

```
zabbixctl -G /guest -a admin
```

## AUTHOR

Egor Kovetskiy <e.kovetskiy@gmail.com>

## CONTRIBUTORS

Lasse Osterild <lasse@oesterild.dk>

Stanislav Seletskiy <s.seletskiy@gmail.com>

Andrey Kitsul <a.kitsul@zarplata.ru>

[GitHub](https://github.com/kovetskiy/zabbixctl)
