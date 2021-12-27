import * as sh from "shelljs";

sh.exec(`
    GOOS=linux go build -o editor
    GOOS=windows go build -o editor
`);
