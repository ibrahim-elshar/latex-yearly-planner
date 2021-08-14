import {copyFileSync, writeFileSync} from "fs";
import {EnvConfig} from "../../common/blocks/envConfig";

const funcs = require('../../common/blocks/funcs');
const t = require('../../common/texblocks/title');
const act = require('../../pages/b5plus/20_annual');
const q = require('../../pages/b5plus/30_quarterly');
const m = require('../../pages/b5plus/40_monthly');
const w = require('../../pages/b5plus/50_weekly');
const d = require('../../pages/b5plus/60_daily');
const td = require('../../pages/b5plus/70_todo');
const nt = require('../../pages/b5plus/80_notes');

['preamble', 'rm2_ddvk', 'macros', 'macros_rm2_ddvk'].forEach(name => {
  copyFileSync(`textpl/${name}.tex`, `out/${name}.tex`);
})

export const buildFiles = (cfg: EnvConfig): void => {
  const {year, weekStart, disableWeeks} = cfg;

  writeFileSync('out/title.tex', t.title(year));
  writeFileSync('out/year.tex', act.annualTable({year, weekStart, weeks: !disableWeeks}))
  writeFileSync('out/quarterlies.tex', funcs.range(1, 5).map(quarter => q.quarter({
    year,
    quarter,
    weekStart,
    weeks: !disableWeeks
  })).join('\n'));
  writeFileSync('out/monthlies.tex', funcs.range(0, 12).map(mn => m.monthlyPage({
    year,
    month: mn,
    weekStart,
    weeks: !disableWeeks
  })).join('\n\\pagebreak\n'))
  disableWeeks ? writeFileSync('out/weeklies.tex', '') : writeFileSync('out/weeklies.tex', w.weeklies(year));
  writeFileSync('out/dailies.tex', d.dailySchedule(year));
  writeFileSync('out/todos.tex', td.todos(year));
  writeFileSync('out/notes.tex', nt.notes(year));
}
