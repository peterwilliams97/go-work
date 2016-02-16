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


if False:
    # lookup vs P

    def func(x, a, b):
        return a * x + b

    p_table = pd.read_csv('p-varies.csv')
    print(p_table)

    x = p_table['P'][9:]
    y = p_table['lookup'][9:]
    print('p_table: %s' % list(p_table))
    print('x: %s:%s %s' % (list(x.shape), x.dtype, [x.min(), x.max()]))
    print('y: %s:%s %s' % (list(y.shape), y.dtype, [y.min(), y.max()]))

    popt, pcov = curve_fit(func, x, y)
    print('popt=%s' % popt)
    print('pcov=%s' % pcov)

    a, b = popt
    print('a=%e,%s' % (a, type(a)))
    print('b=%e,%s' % (b, type(b)))
    xf = np.linspace(0, x.max(), 100)
    yf = func(xf, a, b)
    print('xf: %s:%s %s' % (list(xf.shape), xf.dtype, [xf.min(), xf.max()]))
    print('yf: %s:%s %s' % (list(yf.shape), yf.dtype, [yf.min(), yf.max()]))

    plt.title('lookup = %e * P + %e' % (a, b))
    plt.xlim(0, x.max())
    plt.ylim(0, y.max())
    plt.plot(x, y, label='data')
    plt.plot(xf, yf, label='fit')
    plt.legend()
    plt.show()

if True:
    # index vs T


    def func(x, a, b):
        return a * x + b
        return a * np.log(x + b)

    t_table = pd.read_csv('t-varies.csv')
    print(t_table)
    x = t_table['T']
    y = t_table['lookup']
    print('p_table: %s' % list(t_table))
    print('x: %s:%s %s' % (list(x.shape), x.dtype, [x.min(), x.max()]))
    print('y: %s:%s %s' % (list(y.shape), y.dtype, [y.min(), y.max()]))

    popt, pcov = curve_fit(func, x, y)
    print('popt=%s' % popt)
    print('pcov=%s' % pcov)

    a, b = popt
    print('a=%e,%s' % (a, type(a)))
    print('b=%e,%s' % (b, type(b)))
    xf = np.linspace(0, x.max(), 100)
    yf = func(xf, a, b)
    print('xf: %s:%s %s' % (list(xf.shape), xf.dtype, [xf.min(), xf.max()]))
    print('yf: %s:%s %s' % (list(yf.shape), yf.dtype, [yf.min(), yf.max()]))

    plt.title('index = %e * P + %e' % (a, b))
    plt.xlim(0, x.max())
    plt.ylim(0, y.max())
    plt.plot(x, y, label='data')
    plt.plot(xf, yf, label='fit')
    plt.legend()
    plt.show()

if False:
    # lookup vs T

    t_table = pd.read_csv('t-varies.csv')
    print(t_table)

    x = t_table['T']  #[ 9:]
    y = t_table['lookup']  # [9:]
    x = np.log(x)
    # y = np.log(y)
    print('p_table: %s' % list(p_table))
    print('x: %s:%s %s' % (list(x.shape), x.dtype, [x.min(), x.max()]))
    print('y: %s:%s %s %s' % (list(y.shape), y.dtype, [y.min(), y.max()],
                              [np.argmin(y), np.argmax(y)]))

    def func(x, a, b):
        return np.log(x + b)

    popt, pcov = curve_fit(func, x, y)
    print('popt=%s' % popt)
    print('pcov=%s' % pcov)

    a, b = popt
    print('a=%e,%s' % (a, type(a)))
    print('b=%e,%s' % (b, type(b)))
    xf = np.linspace(0, x.max(), 100)
    yf = func(xf, a, b)
    print('xf: %s:%s %s' % (list(xf.shape), xf.dtype, [xf.min(), xf.max()]))
    print('yf: %s:%s %s' % (list(yf.shape), yf.dtype, [yf.min(), yf.max()]))

    plt.title('lookup = %e * T + %e' % (a, b))
    # plt.xlim(0, x.max())
    # plt.ylim(0, y.max())
    plt.plot(x, y, label='data')
    plt.plot(xf, yf, label='fit')
    plt.legend()
    plt.show()

if False:
    # index vs P
    x = p_table['P']
    y = p_table['index']
    print('p_table: %s' % list(p_table))
    print('x: %s:%s %s' % (list(x.shape), x.dtype, [x.min(), x.max()]))
    print('y: %s:%s %s' % (list(y.shape), y.dtype, [y.min(), y.max()]))

    popt, pcov = curve_fit(func, x, y)
    print('popt=%s' % popt)
    print('pcov=%s' % pcov)

    a, b = popt
    print('a=%e,%s' % (a, type(a)))
    print('b=%e,%s' % (b, type(b)))
    xf = np.linspace(0, x.max(), 100)
    yf = func(xf, a, b)
    print('xf: %s:%s %s' % (list(xf.shape), xf.dtype, [xf.min(), xf.max()]))
    print('yf: %s:%s %s' % (list(yf.shape), yf.dtype, [yf.min(), yf.max()]))

    plt.title('index = %e * P + %e' % (a, b))
    plt.xlim(0, x.max())
    plt.ylim(0, y.max())
    plt.plot(x, y, label='data')
    plt.plot(xf, yf, label='fit')
    plt.legend()
    plt.show()

