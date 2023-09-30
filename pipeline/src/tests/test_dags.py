from pathlib import Path
import pytest
from airflow.models.dagbag import DagBag

dir: Path = Path(__file__).resolve().parent.parent.parent.parent
dataPath: Path = Path(f"{dir}")


# Define una fixture de pytest para crear una instancia de DagBag
@pytest.fixture(params=[f"{dataPath}/pipeline/dags/"])
def dag_bag(request):
    """
    Fixture de pytest para crear una instancia de DagBag para probar DAGs de Airflow.

    Args:
        request: Objeto de solicitud de pytest que contiene el parámetro de prueba.

    Returns:
        DagBag: Una instancia de DagBag para la carpeta de DAG especificada.
    """
    return DagBag(dag_folder=request.param, include_examples=False)


def test_no_import_errors(dag_bag):
    """
    Verifica si no hay errores de importación en las DAGs dentro de un DagBag.

    Esta función realiza una comprobación para asegurarse de que no haya errores de importación
    al cargar las DAGs desde un DagBag. Los errores de importación pueden indicar problemas con
    las dependencias o las importaciones en las DAGs de Airflow.

    Args:
        dag_bag (DagBag): Una instancia de DagBag que contiene las DAGs a verificar.

    Raises:
        AssertionError: Se produce si se encuentran errores de importación en alguna DAG.
    """
    assert not dag_bag.import_errors


def test_requires_tags(dag_bag):
    """
    Verifica que todas las DAGs en un DagBag tengan etiquetas (tags) definidas.

    Esta función itera a través de todas las DAGs en un DagBag y verifica que cada una de ellas tenga al menos
    una etiqueta definida. Las etiquetas son útiles para clasificar y organizar las DAGs en Airflow.

    Args:
        dag_bag (DagBag): Una instancia de DagBag que contiene las DAGs a verificar.

    Raises:
        AssertionError: Se produce si una o más DAGs no tienen etiquetas definidas.
    """
    for dag_id, dag in dag_bag.dags.items():
        assert dag.tags


def test_requires_specific_tag(dag_bag):
    """
    Verifica que todas las DAGs en un DagBag contengan la etiqueta "steam_analytics".

    Esta función itera a través de todas las DAGs en un DagBag y verifica que cada una de ellas contenga
    la etiqueta "steam_analytics". La presencia de esta etiqueta se utiliza para identificar DAGs específicas
    relacionadas con análisis de datos de Steam.

    Args:
        dag_bag (DagBag): Una instancia de DagBag que contiene las DAGs a verificar.

    Raises:
        AssertionError: Se produce si una o más DAGs no contienen la etiqueta "steam_analytics".
    """
    for dag_id, dag in dag_bag.dags.items():
        assert "steam_analytics" in dag.tags


def test_desc_len_greater_than_fifteen(dag_bag):
    """
    Verifica que la longitud de la descripción de todas las DAGs en un DagBag sea mayor que 15 caracteres.

    Esta función itera a través de todas las DAGs en un DagBag y verifica que la longitud de la descripción
    de cada DAG sea mayor que 15 caracteres. La longitud de la descripción es un indicador de la claridad y
    la información proporcionada por la DAG.

    Args:
        dag_bag (DagBag): Una instancia de DagBag que contiene las DAGs a verificar.

    Raises:
        AssertionError: Se produce si la longitud de la descripción de una o más DAGs es menor o igual a 15 caracteres.
    """
    for dag_id, dag in dag_bag.dags.items():
        assert len(dag.description) > 15
