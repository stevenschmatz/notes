\title{Linear Algebra: Vector Dot and Cross Products}
\date{August 18, 2014}
\author{Steven Schmatz, University of Michigan: College of Engineering}

\maketitle
\begin{center} stevenschmatz@gmail.com \end{center} \pagebreak

<!---Content goes here-->

# Vector dot product and vector length

The *dot product*, $\hat{\textbf{a}} \cdot \hat{\textbf{b}}$, is defined as the sum of the product of their corresponding components:

$$\hat{\textbf{a}} \cdot \hat{\textbf{b}} = \sum\limits_{i=1}^n a_ib_i = a_1b_1 + a_2b_2 + a_3b_3 + \dots a_nb_n$$

The vector *magnitude* is defined to be the root mean of squares of all components:

$$\|a\| = \sqrt{a_1^2 + a_2^2 \dots a_n^2}$$

This means that the magnitude can also be written in terms of the dot product:

$$\|a\| = \sqrt{\hat{\textbf{a}} \cdot \hat{\textbf{a}}}$$

# Proving vector dot product properties

The dot product is *commutative*: $\hat{v} \cdot \hat{w} = \hat{w} \cdot \hat{v}$:

$$\hat{v} \cdot \hat{w} = v_1w_1 + v_2w_2 + \dots v_nw_n$$
$$\hat{w} \cdot \hat{v} = w_1v_1 + w_2v_2 + \dots w_nv_n$$
$$v_iw_i = w_iv_i \text{ (commuative property)}$$

Is the dot product *distributive*? Does $(\hat{v} \cdot \hat{w}) \cdot \hat{x} = (\hat{v} \cdot \hat{x} + \hat{w} \cdot \hat{x})$

$$(\hat{v} \cdot \hat{w}) \cdot \hat{x} = (v_1w_1)x_1 + (v_2w_2)x_2 + \dots (v_nw_n)x_n$$
$$(\hat{v} \cdot \hat{x} + \hat{w} \cdot \hat{x}) = v_1x_1 + w_1x_1 + v_2x_2 + w_2x_2 + \dots v_nx_n + w_nx_n$$

You can factor the $x$ term out, yielding:

$$(\hat{v} \cdot \hat{x} + \hat{w} \cdot \hat{x}) = (v_1w_1)x_1 + (v_2w_2)x_2 + \dots (v_nw_n)x_n$$

Hence, dot products are both communative and distributative. Are they also *associative*? Does $(c\hat{v}) \cdot \hat{w} = c(\hat{v} \cdot \hat{w})$?

$$c(\hat{v}) \cdot \hat{w} = cv_1w_1 + cv_2w_2 + \cdots cv_nw_n$$
$$c(\hat{v} \cdot \hat{w}) = cv_1w_1 + cv_2w_2 + \cdots cv_nw_n$$

# Proof of the Cauchy-Schwarz inequality

The Cauchy-Schwarz inequality
  : $\left|\hat{x} \cdot \hat{y}\right| \leq \|x\| \|y\|$, and $\left|\hat{x} \cdot \hat{y}\right| = \|x\| \|y\| \Longleftrightarrow \hat{x} = c\hat{y}$.

For example, say we have a function

$$p(t) = \|t\hat{y} - \hat{x}\|^2 \geq 0$$
$$= (t\hat{y} - \hat{x}) \cdot (t\hat{y} - \hat{x})$$
$$= t\hat{y} \cdot t\hat{y} - 2 \hat{x}\cdot t \hat{y} + \hat{x} \cdot \hat{x}$$
$$= (\hat{y} \cdot \hat{y})t^2 - 2(\hat{x} \cdot \hat{y})t + \hat{x} \cdot \hat{x} \geq 0$$

Let's define $a = (\hat{y} \cdot \hat{y}), b = 2(\hat{x} \cdot \hat{y}), c= \hat{x} \cdot \hat{x}$.

$$p(t) = at^2 - bt + c \geq 0 \text{ for any t}$$
$$p(\frac{b}{2a}) = a\frac{b^2}{4a^2} - b\frac{b}{2a} + c \geq 0$$
$$= \frac{b^2}{4a} - \frac{b^2}{2a} + c = \frac{b^2 - 2b^2}{4a} + c = \frac{-b^2}{4a} + c \geq 0$$
$$c \geq \frac{b^2}{4a}$$
$$4ac \geq b^2$$

Back-substituting $a, b, c$:

$$4(\hat{y} \cdot \hat{y})(\hat{x} \cdot \hat{x}) \leq (2(\hat{x} \cdot \hat{y}))^2$$
$$4(\|\hat{y}\|^2\|\hat{x}\|^2) \leq 2(\hat{x} \cdot \hat{y})\leq 4(\hat{x} \cdot \hat{y})^2$$
$$\|\hat{y}\|^2\|\hat{x}\|^2 \geq (\hat{x} \cdot \hat{y})^2$$
$$\|\hat{y}\|\|\hat{x}\| \geq |\hat{x} \cdot \hat{y}|$$

What if in the case of $\hat{x} = c\hat{y}$?

$$|\hat{x} \cdot \hat{y}| = |c\hat{y} \cdot \hat{y}| = |c||\hat{y} \cdot \hat{y}| = |c|\|\hat{y}\|^2$$
$$= |c|\|\hat{y}\|\|\hat{y}\| = \|c\hat{y}\|\|\hat{y}\|$$
$$= \|\hat{x}\|\|\hat{y}\|$$

# Vector triangle inequality

The *triangle inequality* states that the length of the sum of two vectors is less than or equal to the sum of their magnitudes.

If you have $\hat{x}, \hat{y} \in \mathbb{R}^n, x, y \neq 0$, then $|\hat{x} \cdot \hat{y}| \leq \|\hat{x}\|\|\hat{y}\|$, and if $\hat{x} = c\hat{y} \Longleftrightarrow |\hat{x} \cdot \hat{y}| = \|\hat{x}\|\|\hat{y}\|$.

You could also add $\hat{x} \cdot \hat{y} \leq |\hat{x} \cdot \hat{y}| \leq \|\hat{x}\|\|\hat{y}\|$.

$$\|\hat{x} + \hat{y}\|^2 = (\hat{x} + \hat{y}) \cdot (\hat{x} + \hat{y})$$
$$= \|\hat{x}\|^2 + 2(\hat{x} \cdot \hat{y}) + \|\hat{y}\|^2$$

$$\|\hat{x}+\hat{y}\|^2 \leq \|\hat{x}\|^2 + 2\|\hat{x}\|\|\hat{y}\| + \|\hat{y}\|^2$$
$$\|\hat{x}+\hat{y}\|^2 \leq (\|\hat{x}\| + \|\hat{y}\|)^2$$
$$\|\hat{x}+\hat{y}\| \leq \|\hat{x}\| + \|\hat{y}\|$$

In the case that $\hat{x} = c\hat{y}, c > 0$, then you could say that $\|\hat{x} + \hat{y}\| = \|\hat{x}\| + \|\hat{y}\|$.

# Defining the angle between vectors

If you have vectors $\hat{a}, \hat{b} \in \mathbb{R}^n, \neq 0$, then you could construct a triangle with sides $\|\hat{a}\|, \|\hat{b}\|$, and $\|\hat{a}-\hat{b}\|$. This can work because of the triangle inequality theorem:

$$\|\hat{x} + \hat{y}\| \leq \|\hat{x}\| + \|\hat{y}\|$$

For each of the sides, you can show that you can construct a triangle from nonzero vectors in $\mathbb{R}^n$.

This triangle can allow you to find the *angle* between vectors $\hat{a}$ and $\hat{b}$ using the Law of Cosines, $c^2 = a^2 + b^2 - 2ab\cos(\theta)$:

$$\|\hat{a} - \hat{b}\|^2 = \|b\|^2 + \|a\|^2 - 2\|a\|\|b\|\cos \theta$$
$$(\hat{a}-\hat{b})\cdot(\hat{a}-\hat{b}) = \hat{a}\cdot \hat{a} - 2(\hat{a}\cdot \hat{b}) + \hat{b}\cdot \hat{b}$$
$$(\hat{a} \cdot \hat{b}) = \|\hat{a}\|\|\hat{b}\|\cos \theta$$

This requires two more definitions:

If $\hat{a} = c\hat{b}, c > 0 \Longrightarrow \theta = 0$.
If $\hat{a} = c\hat{b}, c < 0 \Longrightarrow \theta = 180$.

Using this, you can test if two vectors are perpendicular. This is not defined for zero vectors, because you get $\cos \theta = \frac{0}{0}$.

If $\theta = 90^{\circ}$, then $\hat{a}\cdot\hat{b} = 0$. These vectors are *orthogonal*, but not necessarily *perpendicular*.

* Two vectors are **orthogonal** if their dot product is equal to 0. This includes the zero vector, so the zero vector is orthogonal to every vector (even itself!).
* Two vectors are **perpendicular** iff they are orthogonal and *both* nonzero.

# Defining a plane in $\mathbb{R}^3$ with a point and normal vector

The traditional form to write a plane in $\mathbb{R}^3$ would be to write a linear equation $Ax + By + Cz = D$. Another way to write the equation of the plane is by picking a point on the plane, and a vector which is *normal* to all the vectors on the plane. The vector for the point $\hat{x_0}$ itself does not lie on the plane, but it can be used to construct the plane itself.

If you choose another point $\hat{x} = (x, y, z)$, then you know that $(\hat{x} - \hat{x_0})$ has to be on the plane. Hence,

$$\hat{n} \cdot (\hat{x} - \hat{x_0}) = 0$$
$$n_1(x-x_0) + n_2(y-y_0) + n_3(z-z_0) = 0$$

*This is in the form of $Ax + By + Cz = D$!* If you take any point on the plane, you can construct a linear equation for the plane. This is a useful generalization of the plane that can be extended to higher dimensions.

# Cross product introduction

The cross product is useful, but much more limited. The cross product is *only* defined in $\mathbb{R}^3$, and produces a vector from two vectors.

If you have $\hat{a} = (a_1, a_2, a_3), \hat{b} = (b_1, b_2, b_3)$, then $\hat{a} \times \hat{b} = ((a_2b_3 - a_3b_2), (a_3b_1 - a_1b_3), (a_1b_2 - a_2b_1))$.

For example, if you have $\hat{a} = (1, -7, 1), \hat{b} = (5, 2, 4)$, then $\hat{a}\times\hat{b} = (-7*4 - 2, 5 - 4, 2 - (-7 * 5)) = (-30, 1, 37)$. This vector is orthogonal to $\hat{a}$ and $\hat{b}$!

To figure out the direction, use the *right-hand rule*. Point your pointer in the direction of the first vector, your middle finger in the direction of the other, and your thumb points in the direction of the resultant vector.

# Proof: Relationship between cross product and sin of angle

In this section, the proof of $\|\hat{a}\times\hat{b}\| = \|\hat{a}\|\|\hat{b}\|\sin \theta$ is demonstrated.

$$\|\hat{a}\times\hat{b}\|^2 = (a_2b_3 - a_3b_2)^2 + (a_3b_1 - a_1b_3)^2 + (a_1b_2 - a_2b_1)^2$$
$$= a_2^2b_3^2 - 2a_2a_3b_2b_3 + a_3b^2b_2^2 + a_3^2b_1^1 - 2a_1a_3b_1b_3 + a_1^2b_3^2 + a_1^2b_2^2 - 2a_1a_2b_1b_2 + a_2^2b_1^2$$

All the middle terms (with -2 as the coefficient) can be written as $-2(a_2a_3b_2b_3 + a_1a_3b_1b_3 + a_1a_2b_1b_2)$.

We know that $\|\hat{a}\|\|\hat{b}\|\cos\theta = \hat{a}\cdot\hat{b} = a_1b_1 + a_2b_2 + a_3b_3$. Squaring both sides:

$$a_1^2b_1^2 + a_2^2b_2^2 + a_3^2b_3^2 + 2(a_1a_2b_1b_1 + a_1a_3b_1b_3 + a_2a_3b_2b_3)$$

This is equal to the middle term of the $\|\hat{a}\times\hat{b}\|^2$! Hence,

$$\|\hat{a}\times\hat{b}\|^2 + \|\hat{a}\|^2\|\hat{b}\|^2\cos^2\theta = a_1^2(b_1^2 + b_2^2 + b_3^2) + a_2^2(b_1^2 + b_2^2 + b_3^2) + a_3^2(b_1^2 + b_2^2 + b_3^2)$$
$$= (b_1^2 + b_2^2 + b_3^2)(a_1^2 + a_2^2 + a_3^2) = \|\hat{b}\|^2\|\hat{a}\|^2$$
$$\|\hat{a} \times \hat{b}\|^2 = \|\hat{b}\|^2\|\hat{a}\|^2 - \|\hat{b}\|^2\|\hat{a}\|^2 \cos^2\theta $$
$$= \|\hat{a}\|^2\|\hat{b}\|^2(1-\cos^2\theta) = \|\hat{a}\|^2\|\hat{b}\|^2\sin^2\theta$$

Hence,

$$\|\hat{a}\times\hat{b}\| = \|\hat{a}\|\|\hat{b}\|\sin\theta$$

# Dot and cross product comparison/intuition

First, the dot product is defined to be $\hat{a} \cdot \hat{b} = \|\hat{a}\|\|\hat{b}\|\cos\theta$, and the cross product is defined to be $\|\hat{a}\times\hat{b}\| = \|\hat{a}\|\|\hat{b}\|\sin\theta$.

You can imagine the scalar product as being the length of the projection of $\hat{a}$ multiplied by the length of $\hat{b}$. This is why two orthogonal vectors have a dot product of zero, because vector $\hat{a}$ has no projection on $\hat{b}$.

Similar yet opposite, when you compute the cross product of $\hat{a}$ and $\hat{b}$, then $\|\hat{a}\times\hat{b}\|$ can be thought of as the length of $\hat{b}$ multiplied by the component of $\hat{a}$ perpendicular to $\hat{b}$.

The cross product of two vectors is also equal to the *area of the parallelogram* formed by the two vectors!

# Normal vector from plane equation

Earlier in this chapter the method to derive a linear equation for a plane was shown, but how do you find the normal vector from the plane equation $Ax + By + Cz = D$?

Finding the normal vector $\hat{n}$ is incredibly easy - it's $\hat{n} = A\hat{i} + B\hat{j} + C\hat{k}$. The constant $D$ makes no difference, because it would just be shifting the normal vector to a different point on the plane, and not actually affecting the direction of the vector itself.

$D$ is also equal to $ax_p + by_p + cz_p$.

# Point distance to plane

How do you find the minimum distance between a point not on the plane?

If we take the vector between the origin point and the point off the plane, the vector would be:

$$\hat{f} = (x_0 -x_p)\hat{i} + (y_0 - y_p)\hat{j} + (z_0 - z_p)\hat{k}$$

We can use trigonometry to find the minimum distance to the plane by using the *normal vector*. Since the minimum distance from a point $p$ to the plane would be orthogonal to the plane, we can use the vector $f$ and the normal vector to find an angle we can use. After that, you can use trigonometry to find the minimum distance to the plane. The dot product is:

$$\hat{n}\cdot \hat{f} = \frac{Ax_0 - Ax_p + By_0 - By_p + Cz_0 - Cz_p}{\sqrt{A^2 + B^2 + C^2}}$$
$$= \frac{Ax_0 + By_0 + Cz_0 - D}{\sqrt{A^2 + B^2 + C^2}} = \text{ the minimum distance to the plane.}$$

This is the general formula for finding the distance of any point to a plane.