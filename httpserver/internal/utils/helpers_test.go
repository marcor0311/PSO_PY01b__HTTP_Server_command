package utils_test

import (
    "httpserver/internal/constants"
    "httpserver/internal/utils"
    "os"
    "testing"
)

func TestGetEnv(t *testing.T) {
    const key = "TEST_ENV_VAR"

    os.Unsetenv(key)

    def := "default-value"
    if got := utils.GetEnv(key, def); got != def {
        t.Errorf("GetEnv returned %q, want %q when var unset", got, def)
    }

    want := "real-value"
    os.Setenv(key, want)
    if got := utils.GetEnv(key, def); got != want {
        t.Errorf("GetEnv returned %q, want %q when var set", got, want)
    }

    os.Setenv(key, "")
    if got := utils.GetEnv(key, def); got != def {
        t.Errorf("GetEnv returned %q, want %q when var empty", got, def)
    }

    os.Unsetenv(key)
}

func TestIsParallel(t *testing.T) {
    if len(constants.ParallelRoutes) == 0 {
        t.Skip("no parallel routes defined in constants.ParallelRoutes")
    }

    parallelRoute := constants.ParallelRoutes[0]

    if !utils.IsParallel(parallelRoute) {
        t.Errorf("IsParallel(%q) = false, want true", parallelRoute)
    }

    withQuery := parallelRoute + "?foo=bar"
    if !utils.IsParallel(withQuery) {
        t.Errorf("IsParallel(%q) = false, want true", withQuery)
    }

    if utils.IsParallel("/notparallel") {
        t.Errorf("IsParallel(%q) = true, want false", "/notparallel")
    }
}
