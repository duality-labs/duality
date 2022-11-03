import logging

from modelator.pytest.decorators import itf

pytest_plugins = "reactors.bank"


@itf("traces/bank/TestSuccess/violation5.itf.json", keypath="action.tag")
def test_trace():
    logging.info("Successfully executed trace " + "traces/bank/TestSuccess/violation5.itf.json")
