"""
Constants needed for program:
expected bake time
preparation time
"""
EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

def bake_time_remaining(elapsed_bake_time):
    """
    Returns remaining bake time for lasagna.
    This function takes in the number representing the amount of time that has been elapsed in baking the lasagna and calculates the remaining baking time by multiplying it by the expected bake time for the lasagna
    """
    return EXPECTED_BAKE_TIME - elapsed_bake_time

def preparation_time_in_minutes(number_of_layers):
    """
    Return amoount of time needed for preparation of the lasagna.
    This function takes a variable representing the number of layers expected for the lasagna and calculates the time needed for preparing the lasagna by multiplying the number of layers by the preparation time constant
    """
    return number_of_layers * PREPARATION_TIME

def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """
    Return elapsed cooking time.
    This function takes two numbers representing the number of layers & the time already spent baking and calculates the total elapsed minutes spent cooking the lasagna.
    """
    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
