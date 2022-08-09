# `gcr.io/paketo-buildpacks/maven-dist`

The Paketo Maven Dist Buildpack is a Cloud Native Buildpack that builds Maven-based applications from source.

## Behavior

This buildpack will participate all the following conditions are met

* `<APPLICATION_ROOT>/pom.xml` exists or `BP_MAVEN_POM_FILE` is set to an existing POM file.

The buildpack will do the following:

* install either `mvn` or `mvnd` binaries

## Configuration

| Environment Variable        | Description                                                                                                                                                                                                                        |
| --------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `$BP_MAVEN_DAEMON_ENABLED`  | Triggers apache maven-mvnd to be installed and configured for use instead of Maven. The default value is `false`. Set to `true` to use the Maven Daemon.                                                                           |

## Bindings

The buildpack optionally accepts the following bindings:

### Type: `maven`

| Secret                  | Description                                                                                            |
| ----------------------- | ------------------------------------------------------------------------------------------------------ |
| `settings.xml`          | If present `--settings=<path/to/settings.xml>` is prepended to the `maven` arguments                   |
| `settings-security.xml` | If present `-Dsettings.security=<path/to/settings-security.xml>` is prepended to the `maven` arguments |

### Type: `dependency-mapping`

| Key                   | Value   | Description                                                                                       |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------- |
| `<dependency-digest>` | `<uri>` | If needed, the buildpack will fetch the dependency with digest `<dependency-digest>` from `<uri>` |

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0

