package runner

import (
	"path"
	"strings"

	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type BenchmarksManager struct {
	parameters *ParametersManager
	suites     []*BenchmarkSuiteInstance
}

func NewBenchmarksManager(parameters *ParametersManager) *BenchmarksManager {
	c := BenchmarksManager{}
	c.parameters = parameters
	c.suites = make([]*BenchmarkSuiteInstance, 0)
	return &c
}

func (c *BenchmarksManager) Suites() []*BenchmarkSuiteInstance {
	return c.suites
}

func (c *BenchmarksManager) IsSelected() []*BenchmarkInstance {
	benchmarks := make([]*BenchmarkInstance, 0)

	for _, suite := range c.suites {
		for _, benchmark := range suite.benchmarks {
			if benchmark.IsSelected() {
				benchmarks = append(benchmarks, benchmark)
			}
		}
	}
	return benchmarks
}

func (c *BenchmarksManager) SelectAll() {
	for _, suite := range c.suites {
		for _, benchmark := range suite.benchmarks {
			benchmark.Select(true)
		}
	}
}

func (c *BenchmarksManager) SelectByName(benchmarkNames []string) {
	for _, suite := range c.suites {
		for _, benchmark := range suite.benchmarks {
			for _, benchmarkName := range benchmarkNames {
				if benchmarkName == benchmark.FullName() {
					benchmark.Select(true)
				}
			}
		}
	}
}

func (c *BenchmarksManager) Select(benchmarks []*BenchmarkInstance) {
	for _, suite := range c.suites {
		for _, benchmark := range suite.benchmarks {
			for _, anotherBenchmark := range benchmarks {
				if benchmark == anotherBenchmark {
					benchmark.Select(true)
				}
			}
		}
	}
}

func (c *BenchmarksManager) UnselectAll() {
	for _, suite := range c.suites {
		for _, benchmark := range suite.benchmarks {
			benchmark.Select(false)
		}
	}
}

func (c *BenchmarksManager) UnselectByName(benchmarkNames []string) {
	for _, suite := range c.suites {
		for _, benchmark := range suite.benchmarks {
			for _, benchmarkName := range benchmarkNames {
				if benchmarkName == benchmark.FullName() {
					benchmark.Select(false)
				}
			}
		}
	}
}

func (c *BenchmarksManager) Unselect(benchmarks []*BenchmarkInstance) {
	for _, suite := range c.suites {
		for _, benchmark := range suite.benchmarks {
			for _, anotherBenchmark := range benchmarks {
				if benchmark == anotherBenchmark {
					benchmark.Select(false)
				}
			}
		}
	}
}

func (c *BenchmarksManager) AddSuiteFromClass(suiteClassName string) {
	if suiteClassName == "" || len(suiteClassName) == 0 {
		return
	}

	moduleName := suiteClassName
	suiteClassName = ""

	pos := strings.Index(moduleName, ",")
	if pos >= 0 {
		moduleAndClassName := moduleName
		moduleName = moduleAndClassName[0:pos]
		suiteClassName = moduleAndClassName[pos+1:]
	}

	if strings.HasPrefix(moduleName, ".") {
		moduleName = path.Join(moduleName)
	}

	// TODO: Fix dynamic load suit

	// suite := require(moduleName);
	// if suite == nil
	//    panic("Module " + moduleName + " was not found");

	// if (suiteClassName != null && suiteClassName.length > 0)
	//     suite = suite[suiteClassName];

	// if (_.isFunction(suite)) {
	//     suite = new suite();
	//     c.addSuite(suite);
	// }
}

// TODO: Fix dynamic load suit
func (c *BenchmarksManager) AddSuite(suite interface{}) {
	//     if (suite instanceof BenchmarkSuite)
	//         suite = new BenchmarkSuiteInstance(suite);
	//     if (!(suite instanceof BenchmarkSuiteInstance))
	//         throw Error("Incorrect suite type");

	//     c.suites.push(suite);
	//     c.parameters.addSuite(suite);
}

// TODO: Fix dynamic load suit
func (c *BenchmarksManager) AddSuitesFromModule(moduleName string) {
	//     if (moduleName.startsWith("."))
	//         moduleName = path.resolve(moduleName);

	//     let suites = require(moduleName);
	//     if (suites == null)
	//         throw new Error("Module " + moduleName + " was not found");

	//     for (let prop in suites) {
	//         let suite = suites[prop];
	//         if (_.isFunction(suite) && suite.name.endsWith("Suite")) {
	//             try {
	//                 suite = new suite();
	//                 if (suite instanceof BenchmarkSuite) {
	//                     suite = new BenchmarkSuiteInstance(suite);
	//                     c.suites.push(suite);
	//                     c.parameters.addSuite(suite);
	//                 }
	//             } catch (ex) {
	//                 // Ignore
	//             }
	//         }
	//     }
}

func (c *BenchmarksManager) RemoveSuiteByName(suiteName string) {

	var suite *BenchmarkSuiteInstance
	for _, s := range c.suites {
		if s.Name() == suiteName {
			suite = s
		}
	}

	if suite != nil {
		c.parameters.RemoveSuite(suite)
		for index, s := range c.suites {
			if s == suite {
				if index == len(c.suites) {
					c.suites = c.suites[:index-1]
				} else {
					c.suites = append(c.suites[:index], c.suites[index+1:]...)
				}
			}
		}
	}
}

func (c *BenchmarksManager) RemoveSuite(suite interface{}) {
	var localSuite *BenchmarkSuiteInstance = nil
	if curentSuite, ok := suite.(*benchmark.BenchmarkSuite); ok {
		for _, s := range c.suites {
			if curentSuite == s.Suite() {
				localSuite = s
				break
			}
		}
	}
	if localSuite == nil {
		panic("BenchmarksManager:RemoveSuite:Wrong suite type")
	}
	c.parameters.RemoveSuite(localSuite)
	c.RemoveSuiteByName(localSuite.Name())
}

func (c *BenchmarksManager) Clear() {
	for index := 0; index < len(c.suites); index++ {
		suite := c.suites[index]
		c.parameters.RemoveSuite(suite)
	}

	c.suites = make([]*BenchmarkSuiteInstance, 0)
}
