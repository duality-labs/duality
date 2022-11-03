import logging

from modelator.pytest.decorators import itf

pytest_plugins = "reactors.bank"


@itf("traces/bank/TestDuplicate/violation3.itf.json", keypath="action.tag")
def test_trace():
    logging.info("Successfully executed trace " + "traces/bank/TestDuplicate/violation3.itf.json")
