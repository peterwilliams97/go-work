# -*- coding: utf-8 -*-
"""Compute dependencies of golang suffix array indexing and lookup on text and pattern length
T=10000000
dt = a * x + b
a=8.608692e-10
b=4.213161e-05
"""
from __future__ import division, print_function
import numpy as np
import pandas as pd
import matplotlib.pylab as plt
from scipy.optimize import curve_fit

plt.style.use('ggplot')


p_table = pd.read_csv('results-p.csv')
t_table = pd.read_csv('results-t.csv')
print('-' * 80)
print('p_table')
print(p_table)
print('-' * 80)
print('t_table')
print(t_table)


def analyze(title, x, y, func, func_title):
    print('-' * 80)
    print(title)
    print('x: %s:%s %s' % (list(x.shape), x.dtype, [x.min(), x.max()]))
    print('y: %s:%s %s' % (list(y.shape), y.dtype, [y.min(), y.max()]))

    popt, pcov = curve_fit(func, x, y)
    print('popt=%s' % popt)
    print('pcov=\n%s' % pcov)

    a, b = popt
    print('a=%e' % a)
    print('b=%e' % b)
    print(func_title(a, b))
    xf = np.linspace(x.min(), x.max(), 100)
    yf = func(xf, a, b)
    print('xf: %s:%s %s' % (list(xf.shape), xf.dtype, [xf.min(), xf.max()]))
    print('yf: %s:%s %s' % (list(yf.shape), yf.dtype, [yf.min(), yf.max()]))

    plt.title(func_title(a, b))
    # plt.xlim(0, x.max())
    # plt.ylim(0, y.max())
    plt.semilogx(x, y, label='data')
    plt.semilogx(xf, yf, label='fit')
    plt.legend(loc='best')
    plt.savefig('%s.png' % title)
    plt.close()


##########################################################################
# lookup vs P
#
x = p_table['P']
y = p_table['lookup']
T = int(p_table['T'].iloc[0])


def func(x, a, b):
    return a * x + b


def func_title(a, b):
    return 'T=%d. lookup = %e * P + %e' % (T, a, b)

analyze('lookup_vs_P', x, y, func, func_title)


##########################################################################
# lookup vs T
#
x = t_table['T']
y = t_table['lookup']
P = int(t_table['P'].iloc[0])


def func(x, a, b):
    return a * np.log2(x) + b


def func_title(a, b):
    return 'P=%d. lookup = %e log(T) + %e' % (P, a, b)


analyze('lookup_vs_T', x, y, func, func_title)


##########################################################################
# index vs T
#
def func(x, a, b):
    return a * x * np.log2(x) + b


def func_title(a, b):
    return 'lookup = %e T log(T) + %e' % (a, b)

x = t_table['T']
y = t_table['index']
analyze('index_vs_T', x, y, func, func_title)
