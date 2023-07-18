package main

import "os"

var (
	version = "manual build"
	docs    = `zabbixctl ` + version + os.ExpandEnv(`

  zabbixctl is a tool for working with the Zabbix server API using a command line
interface, it provides an effective way for operating on statuses of triggers,
hosts latest data and groups of users.

  zabbixctl reads its configuration from ~/.config/zabbixctl.conf and must be
written using following syntax:

    [server]
      address  = "https://zabbix.local"
      username = "admin"
      password = "password"
      insecure = "false"

    [session]
      path = "~/.cache/zabbixctl.session"

  zabbixctl will authorize against 'zabbix.local' using given user
credentials and save a Zabbix session to ~/.cache/zabbixctl.session.
Subsequent runs will use saved session instead of new re-authorizing.
Zabbix sessions have a default TTL of 15 minutes, so if the saved Zabbix
session is outdated, zabbixctl will repeat authorization and rewrite the
session file.

Usage:
  zabbixctl [options] -T [/<pattern>...]
  zabbixctl [options] -L <hostname>... [/<pattern>...]
  zabbixctl [options] -G [/<pattern>...]
  zabbixctl [options] -M [<hostname>...] [/<pattern>...]
  zabbixctl [options] -H [<pattern>] <hostname>
  zabbixctl -h | --help
  zabbixctl --version

Workflow options:
  -T --triggers
    Search Zabbix trigger statuses. Triggers can be filtered using the
    /<pattern> argument, for example, search and acknowledge all triggers in a
    problem state and match the word 'cache':
      zabbixctl -Tp /cache

    -y --only-nack
      Show only not acknowledged triggers.

    -x --severity
      Specify minimum trigger severity.  Once for information, twice for
      warning, three for disaster, four for high, five for disaster.

    -p --problem
      Show triggers that have a problem state.

    -r --recent
      Show triggers that have recently been in a problem state.

    -s --since <date>
      Show triggers that have changed their state after the given time.
      [default: 7 days ago]

    -u --until <date>
      Show triggers that have changed their state before the given time.

    -m --maintenance
      Show hosts in maintenance.

    -i --sort <fields>
      Show triggers sorted by specified fields.
      [default: lastchange,priority]

    -o --order <order>
      Show triggers in specified order.
      [default: DESC]

    -n --limit <amount>
      Show specified amount of triggers.
      [default: 0]

    -k --acknowledge
      Acknowledge all retrieved triggers.

    -f --noconfirm
      Do not prompt for acknowledge confirmation.

    -d --extended
      Once for printing item's last value from the first component of the
      trigger expression. Twice for adding last value change date. Thrice for
      printing item description as well.

  -L --latest-data
    Search and show latest data for specified host(s). Hosts can be searched for
    using a wildcard character '*'.  Data can be filtered using the /<pattern>
    argument, for example to retrieve latest data for database nodes and search
    information about replication:
      zabbixctl -L dbnode-* /replication

    -g --graph
      Show links on graph pages.

    -w --stacked
      Output single link for the stacked graph of selected data.

    -b --normal
      Output single link for the normal (overlapping) graph of selected data.

  -G --groups
    Search and operate on configuration of usergroups.

    -l --list
     Show list of users in specified usergroup.

    -a --add
     Add specified <user> to specified usergroup.

    -r --remove
     Remove specified <user> from speicifed usergroup.

    -f --noconfirm
     Do not prompt for confirmation.

  -M --maintenances
    Search and operate on configuration of maintenance periods.
    Maintenance can be filtered using the /<pattern> argument, for example,
    search maintenance and match the word 'update-kernel':
      zabbixctl -M dbnode-* /update-kernel

    -a --add <maintenance>
      Add new specified <maintenance> with timeperiod type once.

    --start <date>
      Start date 'yyyy-mm-dd HH:MM'. Default now.

    --end <date>
      Stop date 'yyyy-mm-dd HH:MM'. Default now + period.

    --period <date>
      Period in m/h/d (minutes/hours/days).
      [default: 1d]

    -f --noconfirm
      Do not prompt for confirmation.

    -r --remove <maintenance>
      Remove specified <maintenance>.

    -z --read-stdin
      Read hosts from stdin.

  -H --hosts
    Search and operate on hosts.

    -r --remove <hostname>
      Remove specified <hostname>.


Misc options:
  -c --config <path>
    Use specified configuration file.
    [default: $HOME/.config/zabbixctl.conf]

  -v --verbosity
    Specify program output verbosity.
    Once for debug, twice for trace.

  -h --help
    Show this screen.

  --version
    Show version.
`)
	usage = `
  zabbixctl [options] -T [-v]... [-x]... [-d]... [<pattern>]...
  zabbixctl [options] -L [-v]... <pattern>...
  zabbixctl [options] -G [-v]... [<pattern>]...
  zabbixctl [options] -G [-v]... <pattern>... -a <user>
  zabbixctl [options] -G [-v]... <pattern>... -r <user>
  zabbixctl [options] -M [-v]... [<pattern>]...
  zabbixctl [options] -M [-v]... [<pattern>]... -a <maintenance>
  zabbixctl [options] -M [-v]... -r <maintenance>
  zabbixctl [options] -H [-v]... [<pattern>]...
  zabbixctl [options] -H [-v]... -r <hostname>
  zabbixctl -h | --help
  zabbixctl --version
`
	options = `
Options:
  -T --triggers
    -y --only-nack
    -x --severity
    -p --problem
    -t --recent
    -s --since <date>    [default: 7 days ago]
    -u --until <date>
    -m --maintenance
    -i --sort <fields>   [default: lastchange,priority]
    -o --order <order>   [default: DESC]
    -n --limit <amount>  [default: 0]
    -f --noconfirm
    -k --acknowledge
    -d --extended
  -L --latest-data
    -g --graph
    -w --stacked
    -b --normal
  -G --groups
    -a --add <user>
    -r --remove <user>
  -M --maintenances
    -z --read-stdin
    --period <period>    [default: 1d]
    --start <date>
    --end <date>
  -H --hosts
  -c --config <path>     [default: $HOME/.config/zabbixctl.conf]
  -v --verbosity
  -h --help
  --version
`
)
