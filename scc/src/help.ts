import { bold, command, withSecondary } from "../design/layout";
import { HelpBase } from "@oclif/core";
import { Command } from "@oclif/core/lib/interfaces";
import * as Interfaces from "@oclif/core/lib/interfaces";
import { indentString as indent } from "../tools/strings";
const { say } = require("cfonts");
const { sortBy, uniqBy } = require("../tools/bool");
const { renderList } = require("../tools/list");
const { CommandHelp } = require("../tools/help-command");
const root = require("../contents/helpers/root");
const { TopicFormatter } = require("../contents/helpers/topic");
const { PRIMARY_COLOR } = require("../constants");

function getHelpSubject(args: any) {
  for (const arg of args) {
    if (arg === "--") return;
    if (arg === "help" || arg === "--help" || arg === "-h") continue;
    if (arg.startsWith("-")) return;

    return arg;
  }
}

module.exports = class MyHelpClass extends HelpBase {
  get sortedCommands() {
    let commands = this.config.commands;

    commands = commands.filter((c: any) => this.opts.all || !c.hidden);
    commands = sortBy(commands, (c: any) => c.id);
    commands = uniqBy(commands, (c: any) => c.id);

    console.log(commands);

    return;
  }

  private get _topics(): Interfaces.Topic[] {
    return this.config.topics.filter((topic: Interfaces.Topic) => {
      // it is assumed a topic has a child if it has children
      const hasChild = this.config.topics.some((subTopic) =>
        subTopic.name.includes(`${topic.name}:`)
      );
      return hasChild;
    });
  }

  protected get sortedTopics() {
    let topics = this._topics;
    topics = topics.filter((t: any) => this.opts.all || !t.hidden);
    topics = sortBy(topics, (t: any) => t.name);
    topics = uniqBy(topics, (t: any) => t.name);

    return topics;
  }

  showHelp(args: string[]): Promise<void> {
    // print secman cli version
    console.log(`${bold("Secman Core CLI")} ${this.config.version}`);

    const subject = getHelpSubject(args);

    if (!subject) {
      const rootCmd = this.config.findCommand("");
      if (rootCmd) this.showCommandHelp(rootCmd);
      // this.showRootHelp();
      return this.showRootHelp();
    }

    const command = this.config.findCommand(subject);
    if (command) {
      return this.showCommandHelp(command);
    }

    return this.showRootHelp();
  }

  showCommandHelp(command: Command): Promise<void> {
    const name = command.id;
    const depth = name.split(":").length;

    const subTopics = this.sortedTopics.filter(
      (t: any) =>
        t.name.startsWith(name + ":") && t.name.split(":").length === depth + 1
    );

    const title =
      command.description && this.render(command.description).split("\n")[0];
    if (title) console.log(title + "\n");
    console.log(this.formatCommand(command));
    console.log("");

    if (subTopics.length > 0) {
      console.log(this.formatTopics(subTopics));
      console.log("");
    }

    return root.end();
  }

  async showRootHelp() {
    say("secman", {
      align: "left",
      colors: [PRIMARY_COLOR, "cyan"],
      background: "transparent",
    });

    console.log(command("USAGE\n", true));
    console.log(withSecondary("  $ secman <COMMAND> [FLAGS]\n"));
    this.formatCommands(this.config.commands);
    root.root();
  }

  formatCommand(command: any) {
    const help = new CommandHelp(command, this.config, this.opts);
    return help.generate();
  }

  formatCommands(commands: any) {
    if (commands.length === 0) return "";

    const body = renderList(
      commands.map((c: any) => [
        c.id,
        c.description && this.render(c.description.split("\n")[0]),
      ]),
      {
        spacer: "\n",
        stripAnsi: this.opts.stripAnsi,
        maxWidth: this.opts.maxWidth - 2,
      }
    );

    console.log([command("COMMANDS\n", true), indent(body, 2)].join("\n"));
  }

  formatTopic(topic: any) {
    TopicFormatter(this.render, topic, this.config, this.opts);
  }

  formatTopics(topics: any) {
    if (topics.length === 0) return "";
    const body = renderList(
      topics.map((c: any) => [
        c.name,
        c.description && this.render(c.description.split("\n")[0]),
      ]),
      {
        spacer: "\n",
        stripAnsi: this.opts.stripAnsi,
        maxWidth: this.opts.maxWidth - 2,
      }
    );

    return [command("TOPICS", true), indent(body, 2)].join("\n");
  }
};
