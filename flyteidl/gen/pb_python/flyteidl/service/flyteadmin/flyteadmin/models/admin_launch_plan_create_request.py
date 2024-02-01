# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.admin_launch_plan_spec import AdminLaunchPlanSpec  # noqa: F401,E501
from flyteadmin.models.core_identifier import CoreIdentifier  # noqa: F401,E501


class AdminLaunchPlanCreateRequest(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'id': 'CoreIdentifier',
        'spec': 'AdminLaunchPlanSpec'
    }

    attribute_map = {
        'id': 'id',
        'spec': 'spec'
    }

    def __init__(self, id=None, spec=None):  # noqa: E501
        """AdminLaunchPlanCreateRequest - a model defined in Swagger"""  # noqa: E501

        self._id = None
        self._spec = None
        self.discriminator = None

        if id is not None:
            self.id = id
        if spec is not None:
            self.spec = spec

    @property
    def id(self):
        """Gets the id of this AdminLaunchPlanCreateRequest.  # noqa: E501

        Uniquely identifies a launch plan entity.  # noqa: E501

        :return: The id of this AdminLaunchPlanCreateRequest.  # noqa: E501
        :rtype: CoreIdentifier
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this AdminLaunchPlanCreateRequest.

        Uniquely identifies a launch plan entity.  # noqa: E501

        :param id: The id of this AdminLaunchPlanCreateRequest.  # noqa: E501
        :type: CoreIdentifier
        """

        self._id = id

    @property
    def spec(self):
        """Gets the spec of this AdminLaunchPlanCreateRequest.  # noqa: E501

        User-provided launch plan details, including reference workflow, inputs and other metadata.  # noqa: E501

        :return: The spec of this AdminLaunchPlanCreateRequest.  # noqa: E501
        :rtype: AdminLaunchPlanSpec
        """
        return self._spec

    @spec.setter
    def spec(self, spec):
        """Sets the spec of this AdminLaunchPlanCreateRequest.

        User-provided launch plan details, including reference workflow, inputs and other metadata.  # noqa: E501

        :param spec: The spec of this AdminLaunchPlanCreateRequest.  # noqa: E501
        :type: AdminLaunchPlanSpec
        """

        self._spec = spec

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(AdminLaunchPlanCreateRequest, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, AdminLaunchPlanCreateRequest):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other