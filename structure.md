Trust me, this whole go project structure is confusing to me as well.
Here is how my repo works.

slate/
-   ... go unrelated files
-   internal/ : internal library packages
    -   .../
-   cmd/ : command line tools (main)
    -   slate/
        -   .../
-   pkg/ : public library packages
    -   .../
-   bin/ : compiled binaries
