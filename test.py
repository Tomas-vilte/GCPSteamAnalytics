import locale

initial = 90000

# Establecer la configuración regional para utilizar la coma como separador de miles
locale.setlocale(locale.LC_ALL, 'es_ES.UTF-8')

# Formatear el número con una coma como separador de miles y dos decimales
formatted_price = locale.format_string("ARS$ %.2f", initial / 100, grouping=True)

print(formatted_price)  # Esto imprimirá "ARS$ 82,00"
